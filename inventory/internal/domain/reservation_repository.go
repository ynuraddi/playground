package domain

import "context"

type ReservationRepository interface {
	Save(ctx context.Context, reservation Reservation) error
	Update(ctx context.Context, reservation Reservation) error

	FindById(ctx context.Context, id string) (Reservation, error)
}
