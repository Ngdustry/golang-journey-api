package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// getStatus returns API status check.
func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	status := map[string]string{"status": "ok"}

	json.NewEncoder(w).Encode(status)
	return
}

// New inits new router with endpoints on port
func New() *mux.Router {
	// Route handlers/endpoints
	r := mux.NewRouter()
	tsr := taskSubrouter{}
	usr := userSubrouter{}

	// Base prefix
	api := r.PathPrefix("/api").Subrouter()

	// Health check
	api.HandleFunc("/status", getStatus).Methods("GET")

	// Tasks
	api.HandleFunc("/tasks", tsr.getTasks).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/tasks/{id}", tsr.getTask).Methods(http.MethodGet)
	api.HandleFunc("/tasks", tsr.createTask).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/tasks/update/{id}", tsr.updateTask).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/tasks/delete/{id}", tsr.deleteTask).Methods(http.MethodDelete, http.MethodOptions)

	// Users
	api.HandleFunc("/users", usr.createUser).Methods(http.MethodPost, http.MethodOptions)

	// CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
