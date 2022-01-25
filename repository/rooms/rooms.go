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

func (rr *RoomsRepository) Gets(userId int) ([]entities.Room, error) {
	rooms := []entities.Room{}
	if err := rr.db.Where("user_id=?", userId).Find(&rooms).Error; err != nil {
		return rooms, err
	}
	return rooms, nil
}
func (rr *RoomsRepository) Get(userId int) ([]entities.Room, error) {
	rooms := []entities.Room{}
	if err := rr.db.Not("user_id=?", userId).Find(&rooms).Error; err != nil {

		return rooms, err
	}
	return rooms, nil
}

func (rr *RoomsRepository) Create(newRoom entities.Room) (entities.Room, error) {
	rr.db.Save(&newRoom)
	return newRoom, nil
}

func (rr *RoomsRepository) Update(editRoom entities.Room, roomId int) (entities.Room, error) {
	oldroom := entities.Room{}
	rr.db.Where("user_id=? and id=?", editRoom.User_id, roomId).First(&oldroom)
	rr.db.Model(&oldroom).Updates(editRoom)
	return oldroom, nil
}

func (rr *RoomsRepository) Delete(roomId int, userID uint) (entities.Room, error) {
	room := entities.Room{}
	rr.db.First(&room, "id=? AND user_id=?", roomId, userID).Delete(&room)
	return room, nil

}
