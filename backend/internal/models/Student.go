package models

import (
	"encoding/json"
	"gorm.io/datatypes"
)

type Student struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Percentile  float64  `json:"percentile"`
	Preferences datatypes.JSON `json:"preferences"`
	Allotted    string   `json:"allotted"`
}

func SetPreferences(prefs []string) (datatypes.JSON, error) {
	b, err := json.Marshal(prefs)
	return datatypes.JSON(b), err
}

func GetPreferences(j datatypes.JSON) ([]string, error) {
	var prefs []string
	err := json.Unmarshal(j, &prefs)
	return prefs, err
}