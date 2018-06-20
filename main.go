package main

import (
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatalf("Error loading index page: %s", err)
	}
	t.Execute(w, 0)
}

func init() {

	log.Printf("Initializing Gonsai...")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/bonsais", bonsaisPage)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/bootstrap/css/", http.StripPrefix("/bootstrap/css/", http.FileServer(http.Dir("bootstrap/css"))))
	http.Handle("/bootstrap/js/", http.StripPrefix("/bootstrap/js/", http.FileServer(http.Dir("bootstrap/js"))))

	readSpeciesJson()
	readStylesJson()
}

func main() {
	log.Printf("Starting web server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error initializing web server: %s", err)
	}
}
