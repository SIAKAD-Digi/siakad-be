package request

type CreateOrUpdateClassRequest struct {
	Name string `json:"name" binding:"required,min=5,max=50"`
}
