package database

import (
	"github.com/google/uuid"
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
