package services

import (
	"jeelive/internal/config"
	"jeelive/internal/models"
)

// cutoffs + branches (keep as you had)
var branches = []string{"CS", "IT", "MECH"}

var cutoffs = map[string]float64{
	"CS":   90,
	"IT":   80,
	"MECH": 70,
}

func resolvePreferences(prefs []string) []string {
	if len(prefs) == 1 && prefs[0] == "ALL" {
		return branches
	}
	return prefs
}

// CREATE
func CreateStudent(s models.Student, prefs []string) (models.Student, error) {
	j, err := models.SetPreferences(prefs)
	if err != nil {
		return s, err
	}
	s.Preferences = j
	s.Allotted = ""

	if err := config.DB.Create(&s).Error; err != nil {
		return s, err
	}
	return s, nil
}

// READ
func GetStudents() ([]models.Student, error) {
	var out []models.Student
	if err := config.DB.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

//Show 
func GetStudent(id int) (models.Student, bool) {
	var student models.Student

	result := config.DB.First(&student, id)

	if result.Error != nil {
		return student, false
	}

	return student, true
}
// DELETE
func DeleteStudent(id int) bool {
	res := config.DB.Delete(&models.Student{}, id)
	return res.RowsAffected > 0
}

// CAP
func RunCAP() error {
	var list []models.Student
	if err := config.DB.Find(&list).Error; err != nil {
		return err
	}

	for i := range list {
		prefs, _ := models.GetPreferences(list[i].Preferences)
		resolved := resolvePreferences(prefs)

		best := ""
		bestCut := 0.0

		for _, p := range resolved {
			cut, ok := cutoffs[p]
			if !ok {
				continue
			}
			if list[i].Percentile >= cut && cut > bestCut {
				bestCut = cut
				best = p
			}
		}

		list[i].Allotted = best
		if err := config.DB.Save(&list[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

func EditStudent(id int, name string, percentile float64, prefs []string) (models.Student,bool, error) {
	var s models.Student

	if err := config.DB.First(&s,id).Error; err != nil {
		return s, false, err
	}
	j, err := models.SetPreferences(prefs)
	if err != nil {
		return s, false, err
	}

	s.Name = name
	s.Percentile = percentile
	s.Preferences = j

	s.Allotted = ""

	if err := config.DB.Save(&s).Error; err != nil {
		return s, false, err
	}
	return s, true, nil
}