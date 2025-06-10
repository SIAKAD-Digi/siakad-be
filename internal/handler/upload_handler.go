package handler

import (
	"net/http"
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/response"
	"siakad-digi/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	UploadService *services.UploadService
}

func (h *UploadHandler) UploadImage(ctx *gin.Context) {
	image, err := ctx.FormFile("image")
	maxSize := int64(2 * 1024 * 1024)
	
	if err != nil {
		panic(exception.NewBadRequestError("File tidak ditemukan"))
	}
	
	contenType := image.Header.Get("content-type")

	if image.Size > maxSize {
		panic(exception.NewBadRequestError("File tidak boleh lebih dari 2MB"))
	}

	if !strings.HasPrefix(contenType, "image") {
		panic(exception.NewBadRequestError("File harus gambar"))
	}

	url := h.UploadService.UploadImage(image)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Upload success",
		Data: api.Image{
			ImageUrl: url,
		},
	})
}
