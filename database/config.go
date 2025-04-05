package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// Get database connection string from environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=nanoshell port=5432 sslmode=disable"
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
