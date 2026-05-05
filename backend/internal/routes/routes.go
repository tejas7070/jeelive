package routes

import (
	"github.com/gin-gonic/gin"
	"jeelive/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/students", handlers.GetStudents)
	api.POST("/students", handlers.CreateStudent)
	api.POST("/run-cap", handlers.RunCAP)
	api.GET("/students/:id", handlers.GetStudentByID)
	api.DELETE("/students/:id", handlers.DeleteStudent)
}
