package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID             uint
	User_id        uint `gorm:"not unique"`
	Room_id        uint `gorm:"not unique"`
	Checkin        string
	Checkout       string
	Transaction_id uint
}
