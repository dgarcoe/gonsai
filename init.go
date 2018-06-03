package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Reads the JSON file containing the definition of bonsai species
func readSpeciesJson() {

	jsonSpecies, err := os.Open("resources/species.json")
	if err != nil {
		log.Fatalf("Error reading species JSON file: %s", err)
	}
	defer jsonSpecies.Close()

	byteValue, _ := ioutil.ReadAll(jsonSpecies)
	json.Unmarshal(byteValue, &GonsaiSpecies)

}

//Reads the JSON file containing the definition of bonsai styles
func readStylesJson() {

	jsonStyles, err := os.Open("resources/styles.json")
	if err != nil {
		log.Fatalf("Error reading styles JSON file: %s", err)
	}
	defer jsonStyles.Close()

	byteValue, _ := ioutil.ReadAll(jsonStyles)
	json.Unmarshal(byteValue, &GonsaiStyles)

}
