package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const BONSAIS string = "bonsais"
const EVENTS string = "events"

type GonsaiBonsai struct {
	Id       int
	Name     string
	Age      int
	Species  string
	Style    string
	Acquired string
	Price    float64
	Imgpath  string
	Btype    string
	Events   []GonsaiEvent
}

type GonsaiEvent struct {
	Id      int
	Bonsai  int
	Type    string
	Date    string
	Comment string
}

type BonsaiPageVars struct {
	BonsaiList []GonsaiBonsai
	Species    []string
	Styles     []string
	Types      []string
}

func (b *BonsaiPageVars) setBonsaiList(list []GonsaiBonsai) {
	b.BonsaiList = list
}

func (b *BonsaiPageVars) setBonsaiSpecies(species []string) {
	b.Species = species
}

func (b *BonsaiPageVars) setBonsaiStyles(styles []string) {
	b.Styles = styles
}

func (b *BonsaiPageVars) setBonsaiTypes(types []string) {
	b.Types = types
}

func bonsaisPage(w http.ResponseWriter, r *http.Request) {

	var pageVars BonsaiPageVars

	//Request bonsais to the DB
	bonsaiList, err := getAllBonsaisWithImageAndName("./gonsai.db")
	if err != nil {
		log.Fatalf("Error retrieving bonsai list: %s", err)
	}

	pageVars.setBonsaiList(bonsaiList)
	pageVars.setBonsaiSpecies(GonsaiSpecies)
	pageVars.setBonsaiStyles(GonsaiStyles)
	pageVars.setBonsaiTypes(GonsaiTypes)

	t, err := template.ParseFiles("html/bonsais.html")
	if err != nil {
		log.Fatalf("Error loading bonsais page: %s", err)
	}

	t.Execute(w, pageVars)
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
		err = rows.Scan(&bonsai.Id, &bonsai.Name, &bonsai.Imgpath)
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
