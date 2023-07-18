package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Connnect database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the schema
	// db.AutoMigrate(&Book{})
	err = db.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	// Used to get access to our database
	DB = db
}
