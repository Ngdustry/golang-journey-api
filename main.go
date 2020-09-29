package main

import (
	"log"
	"net/http"

	"github.com/Ngdustry/golang-journey-api/pkg/database"
	"github.com/Ngdustry/golang-journey-api/pkg/router"
)

func main() {
	// Database
	database.New()

	// Router
	r := router.New()
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
