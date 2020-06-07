package main

import (
	"html/template"
	"log"
	"net/http"
)

const POTS string = "pots"

type GonsaiPot struct {
	Id       int
	Name     string
	Pottype  string
	Acquired string
	Price    float64
	Imgpath  string
}

var GonsaiPotList []GonsaiPot

type PotPageVars struct {
	PotList []GonsaiPot
	Types      []string
}

func (b *PotPageVars) setPotList(list []GonsaiPot) {
	b.PotList = list
}

func (b *PotPageVars) setPotTypes(types []string) {
	b.Types = types
}

func potsPage(w http.ResponseWriter, r *http.Request) {

	var pageVars PotPageVars

	//Request pots to the DB
	potList, err := getAllPotsWithImageAndName("./gonsai.db")
	if err != nil {
		log.Fatalf("Error retrieving pot list: %s", err)
	}

	pageVars.setPotList(potList)
	pageVars.setPotTypes(GonsaiPotTypes)

	t, err := template.ParseFiles("html/pots.html")
	if err != nil {
		log.Fatalf("Error loading pots page: %s", err)
	}

	t.Execute(w, pageVars)
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
		err = rows.Scan(&pot.Id, &pot.Name, &pot.Imgpath)
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

