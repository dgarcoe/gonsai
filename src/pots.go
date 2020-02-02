package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const POTS string = "pots"

type GonsaiPot struct {
	id       int
	Name     string
	pottype  string
	acquired float64
	price    float64
	Imgpath  string
}

var GonsaiPotList []GonsaiPot

func potsPage(w http.ResponseWriter, r *http.Request) {

	//Request pots to the DB
	GonsaiPotList, err := getAllPotsWithImageAndName("./gonsai.db")
	if err != nil {
		log.Fatalf("Error retrieving pot list: %s", err)
	}

	t, err := template.ParseFiles("html/pots.html")
	if err != nil {
		log.Fatalf("Error loading pots page: %s", err)
	}

	t.Execute(w, GonsaiPotList)
}

// Returns a list of all pots images and name in the database
func getAllPotsWithImageAndName(databasePath string) ([]GonsaiPot, error) {

	var potlist []GonsaiPot

	db, err := openDatabase(databasePath)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT ID,NAME,IMGPATH from " + POTS)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var pot GonsaiPot
		err = rows.Scan(&pot.id, &pot.Name, &pot.Imgpath)
		if err != nil {
			continue
		}
		potlist = append(potlist, pot)
	}

	rows.Close()

	if err := closeDatabase(db); err != nil {
		return nil, err
	}

	return potlist, nil
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
	err = rows.Scan(&pot.id, &pot.Name, &pot.pottype, &pot.acquired, &pot.price, &pot.Imgpath)
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

	res, err := stmt.Exec(nil, pot.Name, pot.pottype, pot.acquired, pot.price, pot.Imgpath)

	id, err := res.LastInsertId()

	log.Printf("Added new pot with ID: %d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}
