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

type StudentHandler struct {
	StudentService *services.StudentService
}

func (h *StudentHandler) Create(ctx *gin.Context) {
	req := request.CreateStudentRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	h.StudentService.Create(&req)

	ctx.JSON(http.StatusCreated, response.CommonResponse{
		Success: true,
		Message: "Sukses membuat siswa",
	})
}

func (h *StudentHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	req := request.UpdateStudentRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	h.StudentService.Update(&req, id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Sukses mengubah siswa",
	})
}

func (h *StudentHandler) FindAll(ctx *gin.Context) {
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

	req := request.FindAllStudentRequest{
		Name:      ctx.Query("name"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		Status:    ctx.Query("status"),
		Page:      page,
		Limit:     limit,
	}

	studens, total := h.StudentService.FindAll(&req)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Success",
		Data:    studens,
		Meta: api.Pagination{
			Page:    page,
			PerPage: limit,
			Total:   total,
		},
	})
}

func (h *StudentHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	student := h.StudentService.FindById(id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Success",
		Data:    student,
	})
}

func (h *StudentHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	h.StudentService.DeleteById(id)

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Success: true,
		Message: "Sukses menghapus siswa",
	})
}
