package main

const POTS string = "pots"

type GonsaiPot struct {
	id       int
	name     string
	pottype  string
	acquired float64
	price    float64
	imgpath  string
}

var GonsaiPotList []GonsaiPot

// Returns a list of all pots images and name in the database
func getAllPotsWithImageAndName() ([]GonsaiPot, error) {

	var potlist []GonsaiPot

	db, err := openDatabase("gonsai.db")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT ID,NAME,IMGPATH from " + POTS)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var pot GonsaiPot
		err = rows.Scan(&pot.id, &pot.name, &pot.imgpath)
		if err != nil {
			continue
		}
		potlist = append(potlist, pot)
	}

	rows.Close()

	if err := closeDatabase(db); err != nil {
		return nil, err
	}

	return potlist, nil
}
