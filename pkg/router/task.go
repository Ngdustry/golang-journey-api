package router

import (
	"encoding/json"
	"golang-journey-api/pkg/service"
	"net/http"
)

type taskSubrouter struct{}

var ts service.TaskService

func (tsr taskSubrouter) getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", http.MethodGet)
	w.Header().Set("Content-Type", "application/json")

	tasks := ts.GetTasks(r)

	json.NewEncoder(w).Encode(tasks)
	return
}

func (tsr taskSubrouter) createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
	w.Header().Set("Content-Type", "application/json")

	id, err := ts.CreateTask(r)

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
	w.Header().Set("Access-Control-Allow-Methods", http.MethodPut)
	w.Header().Set("Content-Type", "application/json")

	err := ts.UpdateTask(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (tsr taskSubrouter) deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", http.MethodDelete)

	err := ts.DeleteTask(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
