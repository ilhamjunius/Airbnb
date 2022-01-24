package users

import "project-airbnb/entities"

type UsersInterface interface {
	Gets() ([]entities.User, error)
}
