package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type PotInfoPageVars struct {
	Pot       GonsaiPot
}

func potInfo(w http.ResponseWriter, r *http.Request) {

	var pageVars PotInfoPageVars
	var Pot GonsaiPot

	if r.Method == "GET" {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("Error getting ID from URL: %s", err)
		}

		Pot, err = getAllInfoFromPotWithID("./gonsai.db", id)
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

		Pot.Imgpath = "/img/" + handle.Filename

		f, err := os.OpenFile("."+Pot.Imgpath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("Error saving image: %s", err)
		}
		defer f.Close()
		io.Copy(f, file)

		Pot.Name = r.Form["name"][0]
		Pot.Pottype = r.Form["type"][0]
		Pot.Acquired = r.Form["acquired"][0]
		Pot.Price, _ = strconv.ParseFloat(r.Form["price"][0], 64)

		log.Printf("%+v", Pot)
		addNewPot("./gonsai.db", Pot)

	}

	pageVars.Pot = Pot
	t, err := template.ParseFiles("html/potInfo.html")
	if err != nil {
		log.Fatalf("Error loading pots page: %s", err)
	}
	t.Execute(w, pageVars)
}

// Returns all the information from a given pot provided its ID
func getAllInfoFromPotWithID(databasePath string, id int) (GonsaiPot, error) {

	var pot GonsaiPot

	db, err := openDatabase(databasePath)
	if err != nil {
		return pot, err
	}

	rows, err := db.Query("SELECT * from " + POTS + " WHERE ID=" + strconv.Itoa(id))
	if err != nil {
		return pot, err
	}

	rows.Next()
	err = rows.Scan(&pot.Id, &pot.Name, &pot.Pottype, &pot.Acquired, &pot.Price, &pot.Imgpath)
	rows.Close()
	if err != nil {
		return pot, err
	}

	if err := closeDatabase(db); err != nil {
		return pot, err
	}

	return pot, nil

}

// Inserts a new pot in the database
func addNewPot(databasePath string, pot GonsaiPot) error {

	db, err := openDatabase(databasePath)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO " + POTS + " VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(nil, pot.Name, pot.Pottype, pot.Acquired, pot.Price, pot.Imgpath)

	id, err := res.LastInsertId()

	log.Printf("Added new pot with ID: %d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}