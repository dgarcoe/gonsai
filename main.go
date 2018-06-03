package main

import (
	"html/template"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatalf("Error loading index page: %s", err)
	}
	t.Execute(w, 0)
}

func init() {

	log.Printf("Initializing Gonsai...")

	http.HandleFunc("/", test)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	readSpeciesJson()
	readStylesJson()
}

func main() {
	log.Printf("Starting web server")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("Error initializing web server: %s", err)
	}
}
