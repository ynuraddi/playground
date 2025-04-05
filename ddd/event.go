package ddd

import "context"

type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

type EventHandlerFunc func(ctx context.Context, event Event) error

type Event interface {
	EventName() string
}
