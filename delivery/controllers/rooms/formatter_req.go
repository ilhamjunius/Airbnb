package rooms

type AddNewRoomRequestFormat struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Price    int    `json:"price"`
}
type DeleteRoomRequestFormat struct {
	RoomID int `json:"room_id"`
}
