package main

import (
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

func getAllBonsaisWithImageAndName() ([]GonsaiBonsai, error) {

	var bonsailist []GonsaiBonsai

	db, err := openDatabase("gonsai.db")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT ID,NAME,IMGPATH from " + BONSAIS)
	for rows.Next() {
		var bonsai GonsaiBonsai
		err = rows.Scan(&bonsai.id, &bonsai.name, &bonsai.imgpath)
		bonsailist = append(bonsailist, bonsai)
	}

	rows.Close()

	if err := closeDatabase(db); err != nil {
		return nil, err
	}

	return bonsailist, nil
}
