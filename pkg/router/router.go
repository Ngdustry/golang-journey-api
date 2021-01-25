package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"golang-journey-api/pkg/database"
)

// Get status
func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	status := map[string]string{"status": "ok"}

	json.NewEncoder(w).Encode(status)
	return
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	tasks := database.FindAllTasks()

	json.NewEncoder(w).Encode(tasks)
	return
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	task, err := database.FindOneTask(params["id"])

	if err != nil {
		http.Error(w, "Bad Request", 400)
		panic(err)
	} else {
		json.NewEncoder(w).Encode(task)
		return
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	id, err := database.CreateNewTask(r)

	if err != nil {
		w.WriteHeader(422)
		panic(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(id)
		return
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	err := database.UpdateOneTask(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	database.DeleteOneTask(params["id"])

	w.WriteHeader(http.StatusOK)
	return
}

// New inits new router with endpoints on port
func New() *mux.Router {
	// Route handlers/endpoints
	r := mux.NewRouter()

	// Base prefix
	api := r.Host("/api").Subrouter()

	// Health check
	api.HandleFunc("/status", getStatus).Methods("GET")

	// Tasks
	tasks := api.Host("/tasks").Subrouter()
	tasks.HandleFunc("/", getTasks).Methods("GET")
	tasks.HandleFunc("/{id}", getTask).Methods("GET")
	tasks.HandleFunc("/", createTask).Methods("POST", "OPTIONS")
	tasks.HandleFunc("/{id}", updateTask).Methods("PUT", "OPTIONS")
	tasks.HandleFunc("/{id}", deleteTask).Methods("DELETE", "OPTIONS")

	// CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
