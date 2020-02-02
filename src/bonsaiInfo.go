package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type BonsaiInfoPageVars struct {
	Bonsai       GonsaiBonsai
	BonsaiEvents []string
}

func (b *BonsaiInfoPageVars) setBonsaiEvents(events []string) {
	b.BonsaiEvents = events
}

func bonsaiInfo(w http.ResponseWriter, r *http.Request) {

	var pageVars BonsaiInfoPageVars
	var Bonsai GonsaiBonsai

	pageVars.setBonsaiEvents(GonsaiEvents)

	if r.Method == "GET" {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("Error getting ID from URL: %s", err)
		}

		Bonsai, err = getAllInfoFromBonsaiWithID("./gonsai.db", id)
		if err != nil {
			log.Printf("Error retrieving info from database: %s", err)
		}

	} else {

		r.ParseMultipartForm(32 << 20)

		file, handle, err := r.FormFile("image")
		if err != nil {
			log.Printf("Error loading image: %s", err)
		}
		defer file.Close()

		Bonsai.Imgpath = "/img/" + handle.Filename

		f, err := os.OpenFile("."+Bonsai.Imgpath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("Error saving image: %s", err)
		}
		defer f.Close()
		io.Copy(f, file)

		Bonsai.Name = r.Form["name"][0]
		Bonsai.Age, _ = strconv.Atoi(r.Form["age"][0])
		Bonsai.Btype = r.Form["type"][0]
		Bonsai.Species = r.Form["species"][0]
		Bonsai.Style = r.Form["style"][0]
		Bonsai.Acquired = r.Form["acquired"][0]
		Bonsai.Price, _ = strconv.ParseFloat(r.Form["price"][0], 64)

		log.Printf("%+v", Bonsai)
		addNewBonsai("./gonsai.db", Bonsai)
	}

	pageVars.Bonsai = Bonsai
	t, err := template.ParseFiles("html/bonsaiInfo.html")
	if err != nil {
		log.Fatalf("Error loading bonsais page: %s", err)
	}
	t.Execute(w, pageVars)
}

func bonsaiEvent(w http.ResponseWriter, r *http.Request) {

	var pageVars BonsaiInfoPageVars
	var Bonsai GonsaiBonsai
	var Event GonsaiEvent

	pageVars.setBonsaiEvents(GonsaiEvents)
	r.ParseMultipartForm(32 << 20)

	Event.Bonsai, _ = strconv.Atoi(r.Form["bonsaiid"][0])
	Event.Type = r.Form["type"][0]
	Event.Date = r.Form["date"][0]
	Event.Comment = r.Form["comment"][0]

	addNewEvent("./gonsai.db", Event)
	Bonsai, err := getAllInfoFromBonsaiWithID("./gonsai.db", Event.Bonsai)
	if err != nil {
		log.Printf("Error retrieving info from database: %s", err)
	}

	t, err := template.ParseFiles("html/bonsaiInfo.html")
	if err != nil {
		log.Fatalf("Error loading bonsais page: %s", err)
	}

	pageVars.Bonsai = Bonsai
	t.Execute(w, pageVars)
}

// Returns all the information from a given bonsai provided its ID
func getAllInfoFromBonsaiWithID(databasePath string, id int) (GonsaiBonsai, error) {

	var bonsai GonsaiBonsai

	db, err := openDatabase(databasePath)
	if err != nil {
		return bonsai, err
	}

	rows, err := db.Query("SELECT * from " + BONSAIS + " WHERE ID=" + strconv.Itoa(id))
	if err != nil {
		return bonsai, err
	}

	rows.Next()
	err = rows.Scan(&bonsai.Id, &bonsai.Name, &bonsai.Age, &bonsai.Species, &bonsai.Style, &bonsai.Acquired, &bonsai.Price, &bonsai.Imgpath, &bonsai.Btype)
	rows.Close()
	if err != nil {
		return bonsai, err
	}

	rows, err = db.Query("SELECT * from " + EVENTS + " WHERE BONSAIID=" + strconv.Itoa(id))
	if err != nil {
		return bonsai, err
	}

	for rows.Next() {
		var event GonsaiEvent
		err = rows.Scan(&event.Id, &event.Bonsai, &event.Type, &event.Date, &event.Comment)
		if err != nil {
			continue
		}
		bonsai.Events = append(bonsai.Events, event)
	}

	if err := closeDatabase(db); err != nil {
		return bonsai, err
	}

	return bonsai, nil

}

// Inserts a new bonsai in the database
func addNewBonsai(databasePath string, bonsai GonsaiBonsai) error {

	db, err := openDatabase(databasePath)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO " + BONSAIS + " VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(nil, bonsai.Name, bonsai.Age, bonsai.Species, bonsai.Style, bonsai.Acquired, bonsai.Price, bonsai.Imgpath, bonsai.Btype)

	id, err := res.LastInsertId()

	log.Printf("Added new bonsai with ID: %d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}

// Inserts a new event in the database
func addNewEvent(databasePath string, event GonsaiEvent) error {

	db, err := openDatabase(databasePath)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO " + EVENTS + " VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(event.Id, event.Bonsai, event.Type, event.Date, event.Comment)

	id, err := res.LastInsertId()

	log.Printf("Added new event with ID: %d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}
