package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"firstName"`
}

// Task struct
type Task struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Text   string    `json:"text"`
	Status string    `json:"status"`
	User   User      `gorm:"embedded" json:"user"`
}

var db *gorm.DB

// New creates a new database connection.
func New() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PW"), os.Getenv("DBNAME"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	db = database

	// Migrate schema
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&User{})
}

// FindAllTasks will get all tasks.
func FindAllTasks() (res []Task) {
	var tasks []Task
	db.Find(&tasks)

	return tasks
}

// FindOneTask will get specific task.
func FindOneTask(id string) (Task, error) {
	var task Task
	var err error
	result := db.First(&task, id)

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
func DeleteOneTask(id string) {
	db.Where("id = ?", id).Delete(&Task{})
}
