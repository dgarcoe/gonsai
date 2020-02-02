package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func openDatabase(file string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %s", err)
	}
	return db, nil
}

func closeDatabase(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("Error closing database: %s", err)
	}
	return nil
}
