package rooms

type AddNewRoomRequestFormat struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	Price       int    `json:"price"`
	Duration    int    `json:"duration"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
type DeleteRoomRequestFormat struct {
	RoomID int `json:"room_id"`
}
