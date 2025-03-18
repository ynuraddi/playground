package internal

type EventBookingCreate struct {
	SagaId   string `json:"saga_id"`
	UserId   int64  `json:"user_id"`
	RoomId   int64  `json:"room_id"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}

type EventBookingUpdateUserTimeStatus struct {
	Status string `json:"status"`
}
