package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jeelive/internal/models"
	"jeelive/internal/services"
)

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetStudents())
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := services.CreateStudent(student)
	c.JSON(http.StatusCreated, created)
}

func RunCAP(c *gin.Context) {
	services.RunCAP()
	c.JSON(http.StatusOK, gin.H{"message": "CAP round completed"})
}