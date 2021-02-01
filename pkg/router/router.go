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
	api.HandleFunc("/tasks", tsr.getTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", tsr.getTask).Methods("GET")
	api.HandleFunc("/tasks", tsr.createTask).Methods("POST", "OPTIONS")
	api.HandleFunc("/tasks/{id}", tsr.updateTask).Methods("PUT", "OPTIONS")
	api.HandleFunc("/tasks/{id}", tsr.deleteTask).Methods("DELETE", "OPTIONS")

	// Users
	api.HandleFunc("/users", usr.createUser).Methods("POST", "OPTIONS")

	// CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
