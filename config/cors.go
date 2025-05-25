package config

import "github.com/gin-contrib/cors"

func GetCorsConfig() *cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	return &config
} 