package api

import "time"

type TeacherApi struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Nik       string    `json:"nik"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
