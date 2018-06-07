package main

import (
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

func getAllBonsaisWithImageAndName() ([]GonsaiBonsai, error) {

	var bonsailist []GonsaiBonsai

	db, err := openDatabase("gonsai.db")
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

func getAllInfoFromBonsaiWithID(id int) (GonsaiBonsai, error) {

	var bonsai GonsaiBonsai

	db, err := openDatabase("gonsai.db")
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
