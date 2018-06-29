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
	btype    string
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
	err = rows.Scan(&bonsai.id, &bonsai.Name, &bonsai.age, &bonsai.species, &bonsai.style, &bonsai.acquired, &bonsai.price, &bonsai.Imgpath, &bonsai.btype)
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

	stmt, err := db.Prepare("INSERT INTO " + BONSAIS + " VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(nil, bonsai.Name, bonsai.age, bonsai.species, bonsai.style, bonsai.acquired, bonsai.price, bonsai.Imgpath, bonsai.btype)

	id, err := res.LastInsertId()

	log.Printf("Added new bonsai with ID: %d", id)

	if err := closeDatabase(db); err != nil {
		return err
	}

	return nil

}
