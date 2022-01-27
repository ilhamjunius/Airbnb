package books

type NewBookingRequestFormat struct {
	RoomID uint `json:"room_id"`
}
type GetBookingRequestFormat struct {
	RoomID uint `json:"room_id"`
}
type UpdateBookingRequestFormat struct {
	RoomID   uint `json:"room_id"`
	Duration int  `json:"duration"`
}
type CheckoutNowRequestFormat struct {
	BookID uint `json:"book_id"`
}
