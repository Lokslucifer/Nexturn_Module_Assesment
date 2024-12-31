package model

import (
	"time"
)

type Blog struct {
	Id        int       `json:id`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:content`
	Timestamp time.Time `json:timestamp`
}
