package eventpublishing

import "gosanta/internal"

type EventPublisher struct{}

func New() EventPublisher {
	return EventPublisher{}
}

func (e *EventPublisher) PublishEvent(awards.UserAwardEvent) error {
	// STUB TODO
	return nil
}
