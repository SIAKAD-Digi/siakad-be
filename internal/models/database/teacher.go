package database

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	ID        string `gorm:"primaryKey"`
	UsersID   string
	IsMarried bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}
