package rooms

type GetRoomsResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RoomResponse struct {
	ID       uint   `json:"id" form:"id"`
	User_id  uint   `json:"user_id" form:"user_id"`
	Name     string `json:"name" form:"name"`
	Location string `json:"location" form:"location"`
	Price    int    `json:"price" form:"price"`
	Duration int    `json:"duration" form:"duration"`
	Status   string `json:"status" form:"status"`
}
type UserRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
