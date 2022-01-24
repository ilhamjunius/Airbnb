package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID      uint
	Invoice string
	Status  string
	Url     string
	Book    []Book
}
