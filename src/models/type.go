package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name string `json:"name"`
}
