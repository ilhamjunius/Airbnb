package rooms

import (
	"project-airbnb/entities"
	"time"
)

type RoomsInterface interface {
	Create(newRoom entities.Room) (entities.Room, error)
	GetMyRoomIncome(userId int) ([]MyRoomResponseIncome, error)
	Gets(userId int) ([]entities.Room, error)
	GetsById(userId, roomId int) (entities.Room, error)
	Get(userId int) ([]entities.Room, error)
	GetById(userId, roomId int) (entities.Room, error)
	Update(editRoom entities.Room, roomId int) (entities.Room, error)
	Delete(roomID int, userID uint) (entities.Room, error)
}
type MyRoomResponseIncome struct {
	ID       uint      `json:"id" form:"id"`
	User_id  uint      `json:"user_id" form:"user_id"`
	Guest_id uint      `json:"guest_id" form:"guest_id"`
	Book_id  uint      `json:"book_id" form:"book_id"`
	Checkin  time.Time `json:"checkin" form:"checkin"`
	Checkout time.Time `json:"checkout" form:"checkout"`
	Name     string    `json:"name" form:"name"`
	Address  string    `json:"address" form:"address"`
	Location string    `json:"location" form:"location"`
	Price    int       `json:"price" form:"price"`
	Duration int       `json:"duration" form:"duration"`
	Status   string    `json:"status" form:"status"`
}
