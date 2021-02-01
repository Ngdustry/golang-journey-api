package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey" json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

// Task struct
type Task struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Text   string    `json:"text"`
	Status string    `json:"status"`
	UserID string    `json:"userID"`
}
