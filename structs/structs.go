package structs

import "time"

// Substitute struct for holding data
type Substitute struct {
	Date            time.Time `json:"date"`
	Hour            string    `json:"hour"`
	Day             string    `json:"day"`
	Teacher         string    `json:"teacher"`
	TeacherInitials string    `json:"initials"`
	Time            string    `json:"time"`
	Subject         string    `json:"subject"`
	Type            string    `json:"type"`
	Notes           string    `json:"notes"`
	Classes         string    `json:"classes"`
	Room            string    `json:"room"`
	After           string    `json:"after"`
	Cancelled       bool      `json:"cancelled"`
	New             bool      `json:"new"`
	Reason          string    `json:"reason"`
	Counter         string    `json:"counter"`
}

// Version struct for displaying current application versoin
type Version struct {
	Dirty   bool   `json:"dirty"`
	Hash    string `json:"hash"`
	Message string `json:"message"`
}

type SubstituteMeta struct {
	Date     time.Time `json:"date"`
	Class    string    `json:"class"`
	Extended bool      `json:"extended"`
	Updated  time.Time `json:"updated"`
}

type SubstituteResponse struct {
	Substitutes []Substitute   `json:"substitutes"`
	Meta        SubstituteMeta `json:"meta"`
}
