package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (u *UserHandler) GetCurrentUser(ctx *gin.Context) {
	claim, _ := ctx.Get("userClaims")

	ctx.JSON(http.StatusOK, gin.H{
		"user": claim,
	})
}
