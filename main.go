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

	readSpeciesJson()
	readStylesJson()
	readTypesJson()
	readEventsJson()

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/bonsais", bonsaisPage)
	http.HandleFunc("/pots", potsPage)
	http.HandleFunc("/bonsaiInfo", bonsaiInfo)
	http.HandleFunc("/bonsaiEvent", bonsaiEvent)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
}

func main() {
	log.Printf("Starting web server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error initializing web server: %s", err)
	}
}
