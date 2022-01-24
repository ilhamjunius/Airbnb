package rooms

import "project-airbnb/entities"

type RoomsInterface interface {
	Create(newRoom entities.Room) (entities.Room, error)
	Gets() ([]entities.Room, error)
	Update(editRoom entities.Room) (entities.Room, error)
	Delete(roomID int, userID uint) (entities.Room, error)
}
