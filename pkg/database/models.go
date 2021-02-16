package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task struct
type Task struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Text   string    `json:"text"`
	Status string    `json:"status"`
	User   string    `json:"user"`
}
