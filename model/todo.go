package model

import "time"

// Todo Model
type Todo struct {
	Id          string
	UserId      string
	Title       string
	Contents    string
	Start       time.Time
	Due         time.Time
	ActualStart time.Time
	ActualEnd   time.Time
	Status      int
	Version     int
}
