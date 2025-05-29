package request

type CreateTeacherRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=50"`
	Nik         string `json:"nik" binding:"required,len=9,numeric"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,numeric,min=10,max=14"`
	BirthOfDate string `json:"birth_of_date" binding:"required,datetime=2006-01-02"`
	Gender      string `json:"gender" binding:"required,oneof=laki-laki perempuan"`
	Address     string `json:"address" binding:"required,min=25,max=200"`
	IsMarried   *bool  `json:"is_married" binding:"required,boolean"`
}
