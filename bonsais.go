package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const BONSAIS string = "bonsais"

type GonsaiBonsai struct {
	id       int
	Name     string
	age      int
	species  string
	style    string
	acquired float64
	price    float64
	Imgpath  string
}

var GonsaiBonsaiList []GonsaiBonsai

func bonsaisPage(w http.ResponseWriter, r *http.Request) {

	//Request bonsais to the DB
	GonsaiBonsaiList, err := getAllBonsaisWithImageAndName("./gonsai.db")
	if err != nil {
		log.Fatalf("Error retrieving bonsai list: %s", err)
	}

	t, err := template.ParseFiles("html/bonsais.html")
	if err != nil {
		log.Fatalf("Error loading bonsais page: %s", err)
	}

	t.Execute(w, GonsaiBonsaiList)
}

// Returns a list of all bonsais images and name in the database
func getAllBonsaisWithImageAndName(databasePath string) ([]GonsaiBonsai, error) {

	var bonsailist []GonsaiBonsai

	db, err := openDatabase(databasePath)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT ID,NAME,IMGPATH from " + BONSAIS)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var bonsai GonsaiBonsai
		err = rows.Scan(&bonsai.id, &bonsai.Name, &bonsai.Imgpath)
		if err != nil {
			continue
		}
		bonsailist = append(bonsailist, bonsai)
	}

	rows.Close()

	if err := closeDatabase(db); err != nil {
		return nil, err
	}

	return bonsailist, nil
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
	err = rows.Scan(&bonsai.id, &bonsai.Name, &bonsai.age, &bonsai.species, &bonsai.style, &bonsai.acquired, &bonsai.price, &bonsai.Imgpath)
	rows.Close()
	if err != nil {
		return bonsai, err
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

	stmt, err := db.Prepare("INSERT INTO " + BONSAIS + " VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(nil, bonsai.Name, bonsai.age, bonsai.species, bonsai.style, bonsai.acquired, bonsai.price, bonsai.Imgpath)

	id, err := res.LastInsertId()

	log.Printf("Added new bonsai with ID: %d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}
