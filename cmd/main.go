package main

import (
	"log"
	"net/http"

	"golang-journey-api/pkg/database"
	"golang-journey-api/pkg/router"
)

func main() {
	// Database
	database.New()

	// Router
	r := router.New()
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
