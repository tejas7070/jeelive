package models

type Student struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Percentile  float64  `json:"percentile"`
	Preferences []string `json:"preferences"`
	Allotted    string   `json:"allotted"`
}