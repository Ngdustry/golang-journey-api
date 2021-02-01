package router

import (
	"encoding/json"
	"golang-journey-api/pkg/database"
	"net/http"

	"github.com/gorilla/mux"
)

type taskSubrouter struct{}

func (tsr taskSubrouter) getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	tasks := database.FindAllTasks()

	json.NewEncoder(w).Encode(tasks)
	return
}

func (tsr taskSubrouter) getTask(w http.ResponseWriter, r *http.Request) {
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

func (tsr taskSubrouter) createTask(w http.ResponseWriter, r *http.Request) {
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

func (tsr taskSubrouter) updateTask(w http.ResponseWriter, r *http.Request) {
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

func (tsr taskSubrouter) deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	database.DeleteOneTask(params["id"])

	w.WriteHeader(http.StatusOK)
	return
}
