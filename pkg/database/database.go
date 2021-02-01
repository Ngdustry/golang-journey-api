package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
