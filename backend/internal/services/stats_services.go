package services

import (
	"jeelive/internal/config"
	"jeelive/internal/models"
)


type Stats struct {
	Branch string `json:"branch"`
	Total int `json:"total"`
	Filled int `json:"filled"`
	Remaining int `json:"remaining"`
}

func GetStats() ([]Stats, int, int, error) {
	var students []models.Student

	if err := config.DB.Find(&students).Error; err!= nil {
		return nil , 0,0, err
	}

	filled := make(map[string]int)

	for _, s := range students {
		if s.Allotted != "" {
			filled[s.Allotted]++
		}
	}
	var stats []Stats
	totalSeats := 0
	totalFilled:= 0

	for branch, total := range seatLimits {
		f := filled[branch]

		stats =  append(stats, Stats{
			Branch: branch,
			Total: total,
			Filled: f,
			Remaining: total - f,

		})
		totalSeats += total
		totalFilled += f
	}
	return stats, totalSeats, totalFilled, nil
}