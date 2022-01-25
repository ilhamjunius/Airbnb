package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID       uint
	Name     string
	Location string
	Duration int `gorm:"not null"`
	User_id  uint
	Price    int
	Status   string `gorm:"default:Open"`
	Book     []Book
}
