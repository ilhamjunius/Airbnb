package rooms

import (
	"project-airbnb/entities"

	"gorm.io/gorm"
)

type RoomsRepository struct {
	db *gorm.DB
}

func NewRoomsRepo(db *gorm.DB) *RoomsRepository {
	return &RoomsRepository{db: db}
}

func (rr *RoomsRepository) Gets() ([]entities.Room, error) {
	rooms := []entities.Room{}
	rr.db.Find(&rooms)
	return rooms, nil
}

func (rr *RoomsRepository) Create(newRoom entities.Room) (entities.Room, error) {
	rr.db.Save(&newRoom)
	return newRoom, nil
}

func (rr *RoomsRepository) Update(editRoom entities.Room) (entities.Room, error) {
	oldroom := entities.Room{}
	rr.db.Where("user_id=?", editRoom.User_id).Find(&oldroom)

	oldroom.Name = editRoom.Name
	oldroom.Location = editRoom.Location
	oldroom.Price = editRoom.Price

	rr.db.Save(&oldroom)
	return oldroom, nil
}

func (rr *RoomsRepository) Delete(roomId int, userID uint) (entities.Room, error) {
	room := entities.Room{}
	rr.db.Find(&room, "id=? AND user_id=?", roomId, userID).Delete(&room)
	return room, nil

}
