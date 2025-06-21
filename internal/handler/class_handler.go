package handler

import (
	"net/http"
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/request"
	"siakad-digi/internal/models/response"
	"siakad-digi/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	ClassService *services.ClassService
}

func (h *ClassHandler) Create(ctx *gin.Context) {
	req := request.CreateOrUpdateClassRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	h.ClassService.Create(&req)

	ctx.JSON(http.StatusCreated, response.CommonResponse{
		Success: true,
		Message: "Sukses membuat kelas",
	})
}

func (h *ClassHandler) Update(ctx *gin.Context) {
	req := request.CreateOrUpdateClassRequest{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	h.ClassService.Update(&req, id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Sukses mengubah kelas",
	})
}

func (h *ClassHandler) FindAll(ctx *gin.Context) {
	pageQuery := ctx.DefaultQuery("page", "1")
	limitQuery := ctx.DefaultQuery("limit", "10")

	page, errPage := strconv.Atoi(pageQuery)

	if errPage != nil {
		panic(exception.NewBadRequestError("page must number"))
	}

	limit, errSize := strconv.Atoi(limitQuery)

	if errSize != nil {
		panic(exception.NewBadRequestError("size must number"))
	}

	req := request.FindAllClassRequest{
		Name:      ctx.Query("name"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		Page:      page,
		Limit:     limit,
	}

	classes, total := h.ClassService.FindAll(&req)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Success",
		Data:    classes,
		Meta: api.Pagination{
			Page:    page,
			PerPage: limit,
			Total:   total,
		},
	})
}

func (h *ClassHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	h.ClassService.DeleteById(id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Sukses menghapus kelas",
	})
}

func (h *ClassHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	class := h.ClassService.FindById(id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Success",
		Data:    class,
	})
}
