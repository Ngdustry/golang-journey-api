package main

import (
	"log"
	"net/http"
	"os"

	"golang-journey-api/pkg/database"
	"golang-journey-api/pkg/router"
)

func main() {
	// Database
	database.New()

	// Router
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := router.New()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
