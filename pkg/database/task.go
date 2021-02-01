package database

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// FindAllTasks will get all tasks.
func FindAllTasks(userID string) (res []Task) {
	var tasks []Task
	db.Where("user_id = ?", userID).Find(&tasks)

	return tasks
}

// FindOneTask will get specific task.
func FindOneTask(id string) (Task, error) {
	var task Task
	var err error
	result := db.Where("id = ?", id).First(&task)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) == true {
		err = errors.New("No task found")
	}

	return task, err
}

// CreateNewTask will create a new task.
func CreateNewTask(r *http.Request) (id uuid.UUID, err error) {
	var task Task

	json.NewDecoder(r.Body).Decode(&task)

	newID := uuid.New()
	task.ID = newID

	result := db.Create(&task)

	return task.ID, result.Error
}

// UpdateOneTask will update a specific task by ID.
func UpdateOneTask(r *http.Request) (err error) {
	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)
	params := mux.Vars(r)

	result := db.Model(&Task{}).Where("id = ?", params["id"]).Updates(map[string]interface{}{"text": updatedTask.Text, "status": updatedTask.Status})

	return result.Error
}

//DeleteOneTask will delete a specific task by ID.
func DeleteOneTask(id string) (err error) {
	result := db.Where("id = ?", id).Delete(&Task{})

	return result.Error
}
