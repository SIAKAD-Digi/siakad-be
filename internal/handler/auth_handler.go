package handler

import (
	"net/http"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/request"
	"siakad-digi/internal/models/response"
	"siakad-digi/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	req := request.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	jwt := h.AuthService.Login(&req)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Login success",
		Data: api.AuthApi{
			AccessToken: jwt,
		},
	})
}
