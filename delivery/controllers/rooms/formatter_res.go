package rooms

type GetRoomsResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RoomResponse struct {
	ID       uint
	User_id  uint
	Name     string
	Location string
	Price    int
	Duration int
	Status   string
}
type UserRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
