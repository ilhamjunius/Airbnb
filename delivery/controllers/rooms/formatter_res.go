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
	Address  string `json:"address" form:"address"`
	Location string `json:"location" form:"location"`
	Price    int    `json:"price" form:"price"`
	Duration int    `json:"duration" form:"duration"`
	Status   string `json:"status" form:"status"`
}
type RoomResponseDetail struct {
	ID          uint   `json:"id" form:"id"`
	User_id     uint   `json:"user_id" form:"user_id"`
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	Location    string `json:"location" form:"location"`
	Price       int    `json:"price" form:"price"`
	Duration    int    `json:"duration" form:"duration"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}
type UserRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
