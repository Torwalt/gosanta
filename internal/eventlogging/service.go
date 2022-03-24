package eventlogging

import (
	"fmt"
	awards "gosanta/internal"
	"gosanta/internal/ports"
)

type EventLogger struct {
	eventRepo  ports.EventRepository
	eventQueue ports.EventQueue
}

func New(eventRepo ports.EventRepository, eventQueue ports.EventQueue) EventLogger {
	return EventLogger{eventRepo: eventRepo, eventQueue: eventQueue}
}

func (e *EventLogger) Create(event awards.UserPhishingEvent) error {
	err := e.eventRepo.Write(event)
	if err != nil {
		return err
	}
	return nil
}

// Recursivly consume all current messages from the EventQueue and persist them in the EventRepository.
// Errors on individual messages are logged and the message is skipped.
func (e *EventLogger) LogNewEvents() error {
	msgs, err := e.eventQueue.GetNextMessages()
	if err != nil {
		return fmt.Errorf("could not get messages: %v", err)
	}
	if len(msgs) == 0 {
		return nil
	}
	for _, msg := range msgs {
		action := awards.ToPhishingAction(msg.Action)
		pe := awards.UserPhishingEvent{
			UserID:      awards.UserId(msg.UserId),
			Action:      action,
			CreatedAt:   msg.CreatedAt,
			EmailRef:    msg.EmailRef,
			ProcessedAt: nil,
		}
		err := e.Create(pe)
		if err != nil {
			fmt.Printf("could not persist phishing event: %v", err)
			continue
		}
		err = e.eventQueue.DeleteMessage(msg.EventId)
		if err != nil {
			fmt.Print(err)
			continue
		}
	}

	err = e.LogNewEvents()
	if err != nil {
		return err
	}
	return nil
}
