package rooms

type AddNewRoomRequestFormat struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Price    int    `json:"price"`
	Duration int    `json:"duration"`
	Status   string `json:"status"`
}
type DeleteRoomRequestFormat struct {
	RoomID int `json:"room_id"`
}
