package services

import (
	"strings"

	"jeelive/internal/models"
)

var students []models.Student
var idCounter = 1

// Single source of truth
var branches = []string{"CS", "IT", "MECH"}

// Cutoffs
var cutoffs = map[string]float64{
	"CS":   90,
	"IT":   80,
	"MECH": 50,
}

// Create Student
func CreateStudent(student models.Student) models.Student {
	student.ID = idCounter
	idCounter++
	student.Allotted = ""

	students = append(students, student)
	return student
}

// Get Students
func GetStudents() []models.Student {
	return students
}

// Resolve "ALL" preference
func resolvePreferences(prefs []string) []string {
	if len(prefs) == 1 && strings.ToUpper(prefs[0]) == "ALL" {
		return branches
	}
	return prefs
}

// Run CAP Logic
func RunCAP() {
	for i, student := range students {

		// reset previous allocation
		students[i].Allotted = ""
    bestBranch := ""
		highestCutoff := 0.0

		// handle ALL case
		resolvedPrefs := resolvePreferences(student.Preferences)

		for _, pref := range resolvedPrefs {

			// check if branch exists in cutoff map
			cutoff, exists := cutoffs[pref]
			if !exists {
				continue // skip invalid values like "ALL"
			}


			// apply eligibility
			if student.Percentile >= cutoff {
				if cutoff > highestCutoff {
					highestCutoff = cutoff
					bestBranch = pref
				}
			}
		}
		students[i].Allotted = bestBranch
	}
}

func DeleteStudent(id int) bool {
	for i, student := range students {
		if student.ID == id {
			students =  append(students[:i], students[i+1:]...)
			return true 
		}

}
	return false
}