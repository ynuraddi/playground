package internal

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type SagaStatus string

const (
	SagaStatusPending   SagaStatus = "pending"
	SagaStatusCompleted SagaStatus = "completed"

	SagaStatusFailed    SagaStatus = "canceling"
	SagaStatusCancelled SagaStatus = "canceled"
)

type SagaEvent string

const (
	BookingRoom SagaEvent = "booking_room"
)

type Topic string

const (
	BookingCreate Topic = "booking_create"
	BookingCancel Topic = "booking_cancel"

	BookingUpdateUserTimeStatus Topic = "booking_user_time_update_status"
	BookingPaymentCreateStatus         Topic = "booking_payment_create_status"
	BookingRoomReservationStatus       Topic = "booking_room_reservate_status"
)

type service struct {
	eventBus *kafka.Writer
	repo     repository
}

func NewBookingService(eventBus *kafka.Writer, repo repository) *service {
	return &service{eventBus: eventBus, repo: repo}
}

func (s *service) CreateBooking(ctx context.Context, dto CreateBookingDTO) (id int64, err error) {
	dto.SagaId = uuid.NewString()

	return id, s.repo.DoTx(ctx, func(txCtx context.Context) error {
		if err = s.repo.CreateSaga(txCtx, sagaEntry{
			id:     dto.SagaId,
			event:  string(BookingRoom),
			status: string(SagaStatusPending),
		}); err != nil {
			return err
		}

		id, err = s.repo.CreateBooking(txCtx, dto)
		if err != nil {
			return err
		}

		payload, err := json.Marshal(EventBookingCreate(dto))
		if err != nil {
			return err
		}

		return s.eventBus.WriteMessages(txCtx, kafka.Message{
			Topic: string(BookingCreate),
			Value: payload,
		})
	})
}

func (s *service) HandleBookingUpdateUserTimeStatus(value []byte) error {
	var event 
	if err := json.Unmarshal(value, &event); err != nil {
		return err
	}
}
