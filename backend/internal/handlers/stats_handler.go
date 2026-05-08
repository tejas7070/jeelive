package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"jeelive/internal/services"
)

func GetStats(c *gin.Context) {
	stats, totalSeats, totalFilled, err := services.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stats":        stats,
		"total_seats":  totalSeats,
		"total_filled": totalFilled,
	})
}