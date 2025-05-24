package database

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID              string `gorm:"primaryKey"`
	UsersID         string
	StudentGuardian string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DeletedAt       gorm.DeletedAt
}
