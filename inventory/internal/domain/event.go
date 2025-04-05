package domain

type Event interface {
	EventName() string
}

type EventProducer struct {
	events []Event
}

func (ep *EventProducer) AddEvent(event Event) {
	ep.events = append(ep.events, event)
}

func (ep *EventProducer) GetEvents() []Event {
	return ep.events
}
