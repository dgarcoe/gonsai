package main

var GonsaiSpecies []string

var GonsaiStyles []string

type GonsaiBonsai struct {
	id       int
	name     string
	age      int
	species  string
	style    string
	acquired float64
	price    float64
}

var GonsaiBonsaiList []GonsaiBonsai
