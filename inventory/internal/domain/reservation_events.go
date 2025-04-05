package domain

type EventReservationCreated struct {
	Reservation Reservation
}

func (event EventReservationCreated) EventName() string { return "inventory.ReservationCreated" }

type EventReservationCanceled struct {
	Reason error
}

func (event EventReservationCanceled) EventName() string { return "inventory.ReservationCanceled" }

type EventReservationConfirmed struct {
	Reservation Reservation
}

func (event EventReservationConfirmed) EventName() string { return "inventory.ReservationConfirmed" }
