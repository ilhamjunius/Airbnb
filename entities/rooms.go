package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID       uint
	Name     string
	Location string
	Duration int
	User_id  uint
	Price    int
	Status   string
}
