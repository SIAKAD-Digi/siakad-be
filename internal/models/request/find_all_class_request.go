package request

type FindAllClassRequest struct {
	Name      string
	StartDate string
	EndDate   string
	Page      int
	Limit     int
}
