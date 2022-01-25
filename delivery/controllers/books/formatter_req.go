package books

type NewBookingRequestFormat struct {
	RoomID uint `json:"room_id"`
}
type GetBookingRequestFormat struct {
	RoomID uint `json:"room_id"`
}
