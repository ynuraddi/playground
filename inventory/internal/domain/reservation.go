package domain

import (
	"github.com/stackus/errors"
	"github.com/ynuraddi/playground/ddd"
)

var (
	ErrReservationStatusNotAllowed = errors.Wrap(errors.ErrBadRequest, "reservation status is not allowed")
)

type Reservation struct {
	EventProducer
	Id       string
	Products []Product
	Status   ReservationStatus
}

func (r *Reservation) Events() (events []Event) {
	defer func() { r.events = nil }()
	return r.events
}

type ReservationStatus string

const (
	ReservationStatusPending   ReservationStatus = "pending"
	ReservationStatusCanceled  ReservationStatus = "canceled"
	ReservationStatusConfirmed ReservationStatus = "confirmed"
	ReservationStatusShipped   ReservationStatus = "shipped"
)

func (old ReservationStatus) Can(new ReservationStatus) bool {
	switch new {
	case old:
		return true
	case ReservationStatusCanceled, ReservationStatusConfirmed:
		return old == ReservationStatusPending
	default:
		return false
	}
}

func NewReservation(id string, quantity int64, products ...Product) (Reservation, error) {
	r := Reservation{
		Id:       id,
		Products: products,
	}

	var _ ddd.Event

	for _, p := range products {
		if err := p.SubQuantity(quantity); err != nil {
			return r, err
		}
	}

	r.AddEvent(EventReservationCreated{r})
	return r, nil
}

func (r *Reservation) Cancel() error {
	if !r.Status.Can(ReservationStatusCanceled) {
		return ErrReservationStatusNotAllowed
	}

	r.Status = ReservationStatusCanceled
	return nil
}

func (r *Reservation) Confirm() error {
	if !r.Status.Can(ReservationStatusConfirmed) {
		return ErrReservationStatusNotAllowed
	}

	r.Status = ReservationStatusConfirmed
	r.AddEvent(EventReservationConfirmed{*r})
	return nil
}
