package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID       uint
	Name     string
	Location string
	User_id  uint
	Price    int
}
