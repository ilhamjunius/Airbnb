package entities

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID             uint
	User_id        uint      `gorm:"not unique"`
	Room_id        uint      `gorm:"not unique"`
	Checkin        time.Time `gorm:"default:null"`
	Checkout       time.Time `gorm:"default:null"`
	Checkout_early time.Time `gorm:"default:null"`
	Transaction_id uint
}
