package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"beer-reviewer/models"
)

// PopulateBeers populates the beers database
func PopulateBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	models.Beers = append(models.Beers, models.Beer{ID: 1, Name: "Pilsen Callao", Description: "Mejor cerveza del Perú", Country: "Perú"})
	models.Beers = append(models.Beers, models.Beer{ID: 2, Name: "Cristal", Description: "La cerveza más barata Perú", Country: "Perú"})
	fmt.Fprintf(w, "Beers populated")
}

// GetBeers returns all the beers
func GetBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Beers)
}

// CreateBeer creates a beer
func CreateBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var beer models.Beer
	_ = json.NewDecoder(r.Body).Decode(&beer)
	models.Beers = append(models.Beers, beer)
	json.NewEncoder(w).Encode(beer)
}
