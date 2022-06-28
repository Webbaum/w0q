package models

import "time"

type Url struct {
	ID   string    `json:"id" gorm:"primary_key"`
	Url  string    `json:"url"`
	Date time.Time `json:"date"`
	IP   string    `json:"-"`
}
