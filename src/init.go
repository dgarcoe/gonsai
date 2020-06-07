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

//Reads the JSON file containing the definition of types
func readTypesJson() {

	jsonTypes, err := os.Open("resources/types.json")
	if err != nil {
		log.Fatalf("Error reading types JSON file: %s", err)
	}
	defer jsonTypes.Close()

	byteValue, _ := ioutil.ReadAll(jsonTypes)
	json.Unmarshal(byteValue, &GonsaiTypes)

}

//Reads the JSON file containing the definition of pot types
func readPotTypesJson() {

	jsonTypes, err := os.Open("resources/pot_types.json")
	if err != nil {
		log.Fatalf("Error reading types JSON file: %s", err)
	}
	defer jsonTypes.Close()

	byteValue, _ := ioutil.ReadAll(jsonTypes)
	json.Unmarshal(byteValue, &GonsaiPotTypes)
}

//Reads the JSON file containing the definition of events
func readEventsJson() {

	jsonEvents, err := os.Open("resources/events.json")
	if err != nil {
		log.Fatalf("Error reading types JSON file: %s", err)
	}
	defer jsonEvents.Close()

	byteValue, _ := ioutil.ReadAll(jsonEvents)
	json.Unmarshal(byteValue, &GonsaiEvents)

}
