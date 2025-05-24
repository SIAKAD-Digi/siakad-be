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

type TeacherHandler struct {
	TeacherService *services.TeacherService
}

func (h *TeacherHandler) Create(ctx *gin.Context) {
	req := request.CreateTeacherRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	h.TeacherService.Create(&req)

	ctx.JSON(http.StatusCreated, response.CommonResponse{
		Success: true,
		Message: "Sukses membuat guru",
	})
}

func (h *TeacherHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	req := request.UpdateTeacherRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	h.TeacherService.Update(&req, id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Sukses mengubah guru",
	})
}

func (h *TeacherHandler) FindAll(ctx *gin.Context) {
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

	req := request.FindAllTeacherRequest{
		Name:      ctx.Query("name"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		Status:    ctx.Query("status"),
		Page:      page,
		Limit:     limit,
	}

	teachers, total := h.TeacherService.FindAll(&req)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Success",
		Data:    teachers,
		Meta: api.Pagination{
			Page:    page,
			PerPage: limit,
			Total:   total,
		},
	})
}

func (h *TeacherHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	teacher := h.TeacherService.FindById(id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Success",
		Data:    teacher,
	})
}

func (h *TeacherHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	h.TeacherService.DeleteById(id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Sukses menghapus guru",
	})
}
