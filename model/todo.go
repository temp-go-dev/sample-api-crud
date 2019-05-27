package model

import "time"

// Todo Model
type Todo struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	Start       time.Time `json:"start"`
	Due         time.Time `json:"due"`
	ActualStart time.Time `json:"actualStart"`
	ActualEnd   time.Time `json:"actualEnd"`
	Status      int       `json:"status"`
	Version     int       `json:"version"`
}
