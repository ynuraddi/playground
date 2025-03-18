package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) repository {
	return repository{conn: conn}
}

func (r *repository) DoTx(ctx context.Context, f func(txCtx context.Context) error) (err error) {
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	return f(ctx)
}

type sagaEntry struct {
	id     string `db:"id"`
	event  string `db:"event"`
	status string `db:"status"`
}

func (r repository) CreateSaga(ctx context.Context, saga sagaEntry) error {
	query := `INSERT INTO sagas (id, event, status) VALUES ($1, $2, $3)`
	_, err := r.conn.Exec(ctx, query, saga.id, saga.event, saga.status)
	return err
}

func (r repository) UpdateSagaStatus(ctx context.Context, sagaID, status string) error {
	query := `UPDATE sagas SET status = $1 WHERE id = $2`
	_, err := r.conn.Exec(ctx, query, status, sagaID)
	return err
}

func (r repository) GetSaga(ctx context.Context, sagaID string) (saga sagaEntry, err error) {
	query := `SELECT id, event, status FROM sagas WHERE id = $1`
	if err := r.conn.QueryRow(ctx, query, sagaID).
		Scan(
			&saga.id,
			&saga.event,
			&saga.status,
		); err != nil {
		return sagaEntry{}, err
	}
	return saga, nil
}

type bookingEntry struct {
	id     int64  `db:"id"`
	sagaId string `db:"saga_id"`

	userId   int64  `db:"user_id"`
	roomId   int64  `db:"room_id"`
	checkIn  string `db:"check_in"`
	checkOut string `db:"check_out"`
}

func (r repository) CreateBooking(ctx context.Context, dto CreateBookingDTO) (id int64, err error) {
	query := `INSERT INTO bookings (saga_id, user_id, room_id, check_in, check_out) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = r.conn.QueryRow(ctx, query, dto.SagaId, dto.UserId, dto.RoomId, dto.CheckIn, dto.CheckOut).Scan(&id)
	return id, err
}

func (r repository) GetBookingById(ctx context.Context, bookingID int64) (booking bookingEntry, err error) {
	query := `SELECT id, saga_id, user_id, room_id, check_in, check_out FROM bookings WHERE id = $1`
	if err := r.conn.QueryRow(ctx, query, bookingID).
		Scan(
			&booking.id,
			&booking.sagaId,
			&booking.userId,
			&booking.roomId,
			&booking.checkIn,
			&booking.checkOut,
		); err != nil {
		return bookingEntry{}, err
	}
	return booking, nil
}

func (r repository) GetBookingBySagaId(ctx context.Context, sagaId string) (booking bookingEntry, err error) {
	query := `SELECT id, saga_id, user_id, room_id, check_in, check_out FROM bookings WHERE saga_id = $1`
	if err := r.conn.QueryRow(ctx, query, sagaId).
		Scan(
			&booking.id,
			&booking.sagaId,
			&booking.userId,
			&booking.roomId,
			&booking.checkIn,
			&booking.checkOut,
		); err != nil {
		return bookingEntry{}, err
	}
	return booking, nil
}
