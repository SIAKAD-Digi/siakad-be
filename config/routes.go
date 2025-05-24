package config

import (
	"net/http"
	"siakad-digi/internal/handler"
	"siakad-digi/internal/middleware"
	"siakad-digi/internal/repository"
	"siakad-digi/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//-------- repository ----------------
	userRepository := repository.UserRepository{
		DB: DB,
	}

	studentRepository := repository.StudentRepository{
		DB: DB,
	}

	teacherRepository := repository.TeacherRepository{
		DB: DB,
	}
	// -----------------------------------

	//------- service --------------------
	authService := services.AuthService{
		UserRepository: &userRepository,
	}

	studentService := services.StudentService{
		StudentRepository: &studentRepository,
		UserRepository:    &userRepository,
	}

	teacherService := services.TeacherService{
		TeacherRepository: &teacherRepository,
		UserRepository:    &userRepository,
	}
	// -----------------------------------

	//------- handler --------------------
	authHandler := handler.AuthHandler{
		AuthService: &authService,
	}

	userHandler := handler.UserHandler{}

	studentHandler := handler.StudentHandler{
		StudentService: &studentService,
	}

	teacherHandler := handler.TeacherHandler{
		TeacherService: &teacherService,
	}
	// ------------------------------------

	v1.POST("/auth/login", authHandler.Login)

	protected := v1.Group("/", middleware.JWTMiddleware())

	protected.GET("/me", userHandler.GetCurrentUser)

	// ------ student -------------------------
	protected.POST("/students", studentHandler.Create)
	protected.PUT("/students/:id", studentHandler.Update)
	protected.GET("/students/:id", studentHandler.FindById)
	protected.GET("/students", studentHandler.FindAll)
	protected.DELETE("/students/:id", studentHandler.DeleteById)

	// ------ teacher -------------------------
	protected.POST("/teachers", teacherHandler.Create)
	protected.PUT("/teachers/:id", teacherHandler.Update)
	protected.GET("/teachers/:id", teacherHandler.FindById)
	protected.GET("/teachers", teacherHandler.FindAll)
	protected.DELETE("/teachers/:id", teacherHandler.DeleteById)

	r.Run()
}
