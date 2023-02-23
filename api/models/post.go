package models

import (
	"github.com/AlmazDefourten/goapp/models/tag"
	"time"
)

type Post struct {
	Name      string    `json:"name"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	Date      time.Time `json:"date"`
	Where     string    `json:"where"`
	Likes     int       `json:"likes"`
	Reposts   int       `json:"reposts"`
	Views     int       `json:"views"`
	Tags      tag.Tag   `json:"tags"`
}
