package internal

type Saga struct {
	Id     string     `json:"id"`
	Event  SagaEvent  `json:"event"`
	Status SagaStatus `json:"status"`
}

type Booking struct {
	Id       int64  `json:"id"`
	SagaId   string `json:"saga_id"`
	UserId   int64  `json:"user_id"`
	RoomId   int64  `json:"room_id"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}

type CreateBookingDTO struct {
	SagaId string

	UserId   int64  `json:"user_id"`
	RoomId   int64  `json:"room_id"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}
