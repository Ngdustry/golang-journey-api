package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-journey-api/pkg/database"

	"github.com/gorilla/mux"
)

// TaskService serves as the glue between the Task subrouter and database, performing all related operations and data transformation.
type TaskService struct{}

func (ts TaskService) GetTasks(r *http.Request) []database.Task {
	email, _ := r.Context().Value("Email").(string)
	fmt.Println("GetTasks")
	fmt.Println(email)
	tasks := database.FindAllTasks(email)

	return tasks
}

func (ts TaskService) CreateTask(r *http.Request) (interface{}, error) {
	var data database.Task
	email, _ := r.Context().Value("Email").(string)

	json.NewDecoder(r.Body).Decode(&data)
	data.User = email

	id, err := database.CreateNewTask(data)

	return id, err
}

func (ts TaskService) UpdateTask(r *http.Request) error {
	var updatedTask database.Task

	json.NewDecoder(r.Body).Decode(&updatedTask)
	params := mux.Vars(r)

	err := database.UpdateOneTask(params["id"], updatedTask)

	return err
}

func (ts TaskService) DeleteTask(r *http.Request) error {
	params := mux.Vars(r)

	err := database.DeleteOneTask((params["id"]))

	return err
}
