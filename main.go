package main

import (
	"beer-reviewer/controllers"
	"net/http"
)

func main() {
	r := controllers.BeerRouter()

	http.ListenAndServe(":8000", r)
}
