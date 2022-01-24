package users

import "project-airbnb/entities"

type GetUsersResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}
