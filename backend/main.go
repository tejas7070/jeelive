package main

import (
	"jeelive/internal/routes"

	"github.com/gin-gonic/gin"
	"jeelive/internal/middleware"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	routes.RegisterRoutes(r)

  r.Run(":8081")
}
