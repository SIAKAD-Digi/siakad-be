package request

type FindAllStudentRequest struct {
	Name      string
	StartDate string
	EndDate   string
	Status    string
	Page      int
	Limit     int
}
