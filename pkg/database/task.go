package database

import (
	"github.com/google/uuid"
)

// FindAllTasks will get all tasks.
func FindAllTasks(email string) (res []Task) {
	var tasks []Task
	db.Where(&Task{User: email}).Find(&tasks)

	return tasks
}

// CreateNewTask will create a new task.
func CreateNewTask(user string, task Task) (id uuid.UUID, err error) {
	newID := uuid.New()
	task.ID = newID
	task.User = user

	result := db.Create(&task)

	return task.ID, result.Error
}

// UpdateOneTask will update a specific task by ID.
func UpdateOneTask(id string, updatedTask Task) (err error) {
	result := db.Model(&Task{}).Where("id = ?", id).Updates(map[string]interface{}{"text": updatedTask.Text, "status": updatedTask.Status})

	return result.Error
}

//DeleteOneTask will delete a specific task by ID.
func DeleteOneTask(id string) (err error) {
	result := db.Where("id = ?", id).Delete(&Task{})

	return result.Error
}
