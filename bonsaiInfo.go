package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func bonsaiInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//Load information from an already stored bonsai

		t, err := template.ParseFiles("html/bonsaiInfo.html")
		if err != nil {
			log.Fatalf("Error loading bonsais page: %s", err)
		}

		t.Execute(w, 0)

	} else {
		//Store information from a new bonsai and load it
		var bonsai GonsaiBonsai

		r.ParseMultipartForm(32 << 20)

		file, handle, err := r.FormFile("image")
		if err != nil {
			log.Printf("Error loading image: %s", err)
		}
		defer file.Close()

		bonsai.Imgpath = "/img/" + handle.Filename

		f, err := os.OpenFile("."+bonsai.Imgpath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("Error saving image: %s", err)
		}
		defer f.Close()
		io.Copy(f, file)

		bonsai.Name = r.Form["name"][0]
		bonsai.age, _ = strconv.Atoi(r.Form["age"][0])
		bonsai.btype = r.Form["type"][0]
		bonsai.species = r.Form["species"][0]
		bonsai.style = r.Form["style"][0]
		bonsai.acquired, _ = strconv.ParseFloat(r.Form["acquired"][0], 64)
		bonsai.price, _ = strconv.ParseFloat(r.Form["price"][0], 64)

		log.Printf("%+v", bonsai)

		t, err := template.ParseFiles("html/bonsaiInfo.html")
		if err != nil {
			log.Fatalf("Error loading bonsais page: %s", err)
		}

		t.Execute(w, 0)
	}
}
