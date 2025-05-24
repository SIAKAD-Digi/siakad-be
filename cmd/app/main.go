package main

import (
	"siakad-digi/config"
	"siakad-digi/internal/handler"
	"siakad-digi/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitEnv()
	config.InitLogger()
	config.ConnectDatabase()

	r := gin.Default()

	r.Use(cors.Default())
	r.Use(handler.ErrorHandler(config.Logger))
	r.Use(middleware.NewLoggerMiddleware(config.Logger))

	config.RegisterRoutes(r)

}
