package database

import (
	"encoding/json"
	"net/http"
)

// CreateNewUser will create a new user.
func CreateNewUser(r *http.Request) (id string, err error) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	result := db.Create(&user)

	return user.ID, result.Error
}
