package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey,column:id"`
	Name        string
	Nik         string
	Email       string
	PhoneNumber string
	Password    string
	BirthOfDate string
	Gender      string
	Address     string
	IsActive    bool
	RoleId      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt
}
