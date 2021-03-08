package service

import (
	"encoding/json"
	"net/http"

	"golang-journey-api/pkg/database"

	"github.com/gorilla/mux"
)

// TaskService serves as the glue between the Task subrouter and database, performing all related operations and data transformation.
type TaskService struct{}

// GetTasks returns tasks belonging to user's email.
func (ts TaskService) GetTasks(r *http.Request) []database.Task {
	email, _ := r.Context().Value("Email").(string)
	tasks := database.FindAllTasks(email)

	return tasks
}

// CreateTasks composes new task linked to user's email.
func (ts TaskService) CreateTask(r *http.Request) (interface{}, error) {
	var data database.Task
	email, _ := r.Context().Value("Email").(string)

	json.NewDecoder(r.Body).Decode(&data)
	data.User = email

	id, err := database.CreateNewTask(email, data)

	return id, err
}

// UpdateTask edits a single task belonging to user.
func (ts TaskService) UpdateTask(r *http.Request) error {
	var updatedTask database.Task

	json.NewDecoder(r.Body).Decode(&updatedTask)
	params := mux.Vars(r)

	err := database.UpdateOneTask(params["id"], updatedTask)

	return err
}

// DeleteTask destroys a single task belonging to user.
func (ts TaskService) DeleteTask(r *http.Request) error {
	params := mux.Vars(r)

	err := database.DeleteOneTask((params["id"]))

	return err
}
