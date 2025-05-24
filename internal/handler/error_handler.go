package handler

import (
	"fmt"
	"net/http"
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/response"
	"siakad-digi/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			r := recover()

			if r == nil {
				return
			}

			if err, _ := r.(error); err.Error() == "EOF" {
				logger.Warn("invalid user input", zap.Error(err))
				requestBodyEmptyError(ctx)
				ctx.Abort()
				return
			}

			if err, ok := r.(validator.ValidationErrors); ok {
				logger.Warn("invalid user input", zap.Error(err))
				validationError(ctx, err)
				ctx.Abort()
				return
			}

			if err, ok := r.(*exception.NotFoundError); ok {
				logger.Warn("not found data", zap.Error(err))
				notFoundError(ctx, err)
				ctx.Abort()
				return
			}

			if err, ok := r.(*exception.BadRequestError); ok {
				logger.Warn("invalid user input", zap.Error(err))
				badRequestError(ctx, err)
				ctx.Abort()
				return
			}

			if err, ok := r.(*exception.UnauthorizeError); ok {
				logger.Warn("unauthorize user", zap.Error(err))
				unauthorizeError(ctx, err)
				ctx.Abort()
				return
			}

			logger.Error("something when wrong!", zap.Error(r.(error)))
			internalServerError(ctx)
			ctx.Abort()
		}()
		ctx.Next()
	}
}

func validationError(ctx *gin.Context, errs validator.ValidationErrors) {
	errorMessages := make(map[string]string)

	for _, fe := range errs {
		field := utils.ToSnakeCase(fe.Field())
		errorMessages[field] = validationMessage(fe)
	}

	ctx.JSON(http.StatusBadRequest, response.CommonResponse{
		Success: false,
		Message: "Invalid request body",
		Errors:  errorMessages,
	})
}

func requestBodyEmptyError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, response.CommonResponse{
		Success: false,
		Message: "Request body can't empty",
	})
}

func notFoundError(ctx *gin.Context, err *exception.NotFoundError) {
	ctx.JSON(http.StatusNotFound, response.CommonResponse{
		Success: false,
		Message: err.Error(),
	})
}

func badRequestError(ctx *gin.Context, err *exception.BadRequestError) {
	ctx.JSON(http.StatusBadRequest, response.CommonResponse{
		Success: false,
		Message: err.Error(),
	})
}

func unauthorizeError(ctx *gin.Context, err *exception.UnauthorizeError) {
	ctx.JSON(http.StatusUnauthorized, response.CommonResponse{
		Success: false,
		Message: err.Error(),
	})
}

func internalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, response.CommonResponse{
		Success: false,
		Message: "Internal Server Error",
	})
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "tidak boleh kosong"
	case "email":
		return "format harus email"
	case "len":
		return fmt.Sprintf("harus %s karakter", fe.Param())
	case "min":
		return fmt.Sprintf("minimal %s karaker", fe.Param())
	case "max":
		return fmt.Sprintf("maximal %s karaker", fe.Param())
	case "datetime":
		return fmt.Sprintf("format harus %s", fe.Param())
	case "numeric":
		return "karakter harus berupa angka"
	case "boolean":
		return "harus true atau false"
	default:
		return fmt.Sprintf("tag %s not register", fe.Tag())
	}
}
