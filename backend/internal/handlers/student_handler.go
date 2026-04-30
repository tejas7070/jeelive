package handlers

import (
	"net/http"
	"strconv"

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

	if student.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	if student.Percentile < 0 || student.Percentile > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Percentile must be between 0 and 100"})
		return
	}

	if len(student.Preferences) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one preference is required"})
		return
	}

	created := services.CreateStudent(student)
	c.JSON(http.StatusCreated, created)
}

func RunCAP(c *gin.Context) {
	services.RunCAP()
	c.JSON(http.StatusOK, gin.H{"message": "CAP round completed"})
}

func DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deleted := services.DeleteStudent(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
