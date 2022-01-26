package rooms

import "project-airbnb/entities"

type RoomsInterface interface {
	Create(newRoom entities.Room) (entities.Room, error)
	Gets(userId int) ([]entities.Room, error)
	GetsById(userId, roomId int) (entities.Room, error)
	Get(userId int) ([]entities.Room, error)
	GetById(userId, roomId int) (entities.Room, error)
	Update(editRoom entities.Room, roomId int) (entities.Room, error)
	Delete(roomID int, userID uint) (entities.Room, error)
}
