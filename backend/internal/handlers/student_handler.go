package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"jeelive/internal/models"
	"jeelive/internal/services"
)

type createReq struct {
	Name        string   `json:"name"`
	Percentile  float64  `json:"percentile"`
	Preferences []string `json:"preferences"`
}

// POST /students
func CreateStudent(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	// validations
	if req.Name == "" {
		c.JSON(400, gin.H{"error": "Name required"})
		return
	}
	if req.Percentile < 0 || req.Percentile > 100 {
		c.JSON(400, gin.H{"error": "Percentile 0–100"})
		return
	}
	if len(req.Preferences) == 0 {
		c.JSON(400, gin.H{"error": "Select preferences"})
		return
	}

	s, err := services.CreateStudent(models.Student{
		Name:       req.Name,
		Percentile: req.Percentile,
	}, req.Preferences)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// convert JSON -> []string for response
	prefs, _ := models.GetPreferences(s.Preferences)

	c.JSON(201, gin.H{
		"id":          s.ID,
		"name":        s.Name,
		"percentile":  s.Percentile,
		"preferences": prefs,
		"allotted":    s.Allotted,
	})
}

// GET /students
func GetStudents(c *gin.Context) {
	list, err := services.GetStudents()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resp []gin.H
	for _, s := range list {
		prefs, _ := models.GetPreferences(s.Preferences)
		resp = append(resp, gin.H{
			"id":          s.ID,
			"name":        s.Name,
			"percentile":  s.Percentile,
			"preferences": prefs,
			"allotted":    s.Allotted,
		})
	}

	c.JSON(200, resp)
}

// DELETE /students/:id
func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	if !services.DeleteStudent(id) {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"message": "deleted"})
}

// POST /run-cap
func RunCAP(c *gin.Context) {
	if err := services.RunCAP(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "CAP completed"})
}