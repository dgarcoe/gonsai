package main

import (
	"database/sql"
	"fmt"
)

func openDatabase(file string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./gonsai.db")
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %s", err)
	}
	return db, nil
}

func getBonsais() ([]GonsaiBonsai, error) {

	db, err := openDatabase("./gonsai.db")
	if err != nil {
		return nil, err
	}

	//TODO Use bonsai table as a constant in other place
	rows, err := db.Query("SELECT * from bonsais")
	for rows.Next() {

	}

	rows.Close()

	db.Close()

	return nil, nil
}
