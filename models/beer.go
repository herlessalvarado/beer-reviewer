package models

// Beer is the inital beer definition
type Beer struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

// Beers is a slice of beers
var Beers []Beer
