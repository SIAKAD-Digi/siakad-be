package api

import "time"

type StudentApi struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Nik       string    `json:"nik"`
	ClassName *string   `json:"class_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
