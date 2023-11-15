package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	PokedexNumber  uint       `json:"pokedex_number"`
	Name           string     `json:"name"`
	FirstTypeId    uint       `json:"-"`
	SecondTypeId   uint       `json:"-"`
	Total          int        `json:"total"`
	HP             uint       `json:"hp"`
	Attack         uint       `json:"attack"`
	Defense        uint       `json:"defense"`
	SpecialAttack  uint       `json:"special_attack"`
	SpecialDefense uint       `json:"special_defense"`
	Speed          uint       `json:"speed"`
	GenerationId   uint       `json:"-"`
	Legendary      bool       `json:"legendary"`
	Mega           bool       `json:"mega"`
	FirstType      Type       `gorm:"foreignKey:FirstTypeId" json:"firstType"`
	SecondType     Type       `gorm:"foreignKey:SecondTypeId" json:"secondType"`
	Generation     Generation `gorm:"foreignKey:GenerationId" json:"generation"`
}
