package main

import (
	"github.com/gin-gonic/gin"
	"jeelive/internal/config"
	"jeelive/internal/middleware"
	"jeelive/internal/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	routes.RegisterRoutes(r)

	r.Run(":8080")
}