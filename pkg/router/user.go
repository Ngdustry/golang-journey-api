package router

import (
	"encoding/json"
	"golang-journey-api/pkg/database"
	"net/http"
)

type userSubrouter struct{}

func (usr userSubrouter) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	id, err := database.CreateNewUser(r)

	if err != nil {
		w.WriteHeader(422)
		panic(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(id)
		return
	}
}
