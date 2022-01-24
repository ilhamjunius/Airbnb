package users

import (
	"project-airbnb/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Gets() ([]entities.User, error) {
	users := []entities.User{}
	ur.db.Find(&users)
	return users, nil
}
