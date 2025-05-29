package api

type TeacherDetailApi struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Nik            string  `json:"nik"`
	PhoneNumber    string  `json:"phone_number"`
	Email          string  `json:"email"`
	ProfilePicture *string `json:"profile_picture"`
	BirthOfDate    string  `json:"birth_of_date"`
	Gender         string  `json:"gender"`
	Address        string  `json:"address"`
	IsActive       bool    `json:"is_active"`
	IsMarried      bool    `json:"is_married"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
