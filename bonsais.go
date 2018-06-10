package main

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const BONSAIS string = "bonsais"

type GonsaiBonsai struct {
	id       int
	name     string
	age      int
	species  string
	style    string
	acquired float64
	price    float64
	imgpath  string
}

var GonsaiBonsaiList []GonsaiBonsai

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
		err = rows.Scan(&bonsai.id, &bonsai.name, &bonsai.imgpath)
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
	err = rows.Scan(&bonsai.id, &bonsai.name, &bonsai.age, &bonsai.species, &bonsai.style, &bonsai.acquired, &bonsai.price, &bonsai.imgpath)
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

	res, err := stmt.Exec(nil, bonsai.name, bonsai.age, bonsai.species, bonsai.style, bonsai.acquired, bonsai.price, bonsai.imgpath)

	id, err := res.LastInsertId()

	log.Printf("%d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}
