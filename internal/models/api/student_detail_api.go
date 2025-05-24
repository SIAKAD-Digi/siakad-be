package api

type StudentDetailApi struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Nik             string  `json:"nik"`
	PhoneNumber     string  `json:"phone_number"`
	Email           string  `json:"email"`
	ClassName       *string `json:"class_name"`
	ProfilePicture  *string `json:"profile_picture"`
	BirthOfDate     string  `json:"birth_of_date"`
	Address         string  `json:"address"`
	IsActive        bool    `json:"is_active"`
	StudentGuardian string  `json:"student_guardian"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
