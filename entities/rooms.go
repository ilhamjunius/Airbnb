package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID         uint
	Name       string
	Address    string
	Location   string
	Duration   int `gorm:"not null"`
	User_id    uint
	Price      int
	Status     string `gorm:"default:OPEN"`
	Desciption string
	Book       []Book
}
