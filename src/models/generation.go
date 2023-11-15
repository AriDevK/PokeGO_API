package models

import "gorm.io/gorm"

type Generation struct {
	gorm.Model
	Number int    `json:"number" xml:"number"`
	Name   string `json:"name" xml:"name"`
}
