package main

import (
	"log"

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
}

var GonsaiBonsaiList []GonsaiBonsai

func getAllBonsais() ([]GonsaiBonsai, error) {

	db, err := openDatabase("gonsai.db")
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	rows, err := db.Query("SELECT * from " + BONSAIS)
	for rows.Next() {
		var bonsai GonsaiBonsai
		err = rows.Scan(&bonsai.id, &bonsai.name, &bonsai.age, &bonsai.species, &bonsai.style, &bonsai.acquired, &bonsai.price)
		log.Printf("%+v", bonsai)
	}

	rows.Close()

	if err := closeDatabase(db); err != nil {
		return nil, err
	}

	return nil, nil
}
