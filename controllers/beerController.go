package controllers

import (
	"beer-reviewer/services"

	"github.com/gorilla/mux"
)

// BeerRouter handles all controllers related to beers
func BeerRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/beers", services.PopulateBeers).Methods("POST")
	r.HandleFunc("/beers", services.GetBeers).Methods("GET")
	r.HandleFunc("/beer", services.CreateBeer).Methods("POST")

	return r
}
