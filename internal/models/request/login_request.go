package request

type LoginRequest struct {
	Nik      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}
