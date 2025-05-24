package request

type FindAllTeacherRequest struct {
	Name      string
	StartDate string
	EndDate   string
	Status    string
	Page      int
	Limit     int
}
