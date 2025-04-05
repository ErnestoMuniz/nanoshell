package models

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primarykey;size:8" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"size:255;not null;unique" json:"username"`
	Email     string         `gorm:"size:255;not null;unique" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	Admin     bool           `gorm:"not null" json:"admin"`
	Active    bool           `gorm:"default:true" json:"active"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := gonanoid.New(8)
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}
