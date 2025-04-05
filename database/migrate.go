package database

import (
	"nanoshell/database/models"
	"nanoshell/utils"

	"gorm.io/gorm"
)

func AutoMigrate() error {
	if err := DB.AutoMigrate(
		&models.User{},
	); err != nil {
		return err
	}

	// Check if admin user exists
	var adminUser models.User
	result := DB.Where("username = ?", "admin").First(&adminUser)

	// If admin user doesn't exist, create it
	if result.Error == gorm.ErrRecordNotFound {
		adminUser = models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: utils.HashPassword("admin123"), // This should be hashed in a real application
			Admin:    true,
			Active:   true,
		}
		if err := DB.Create(&adminUser).Error; err != nil {
			return err
		}
	}

	return nil
}
