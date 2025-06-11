package eventlogging_test

import (
	"testing"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/eventlogging"
	"gosanta/internal/mocks"
	events "gosanta/pkg"

	"github.com/golang/mock/gomock"
)

func TestLogNewEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockEventRepository(ctrl)
	mockQueue := mocks.NewMockEventQueue(ctrl)

	eventID := "123-321"
	expMsg := events.PhishingEvent{
		EventId:   eventID,
		UserId:    1,
		Action:    "opened",
		CreatedAt: time.Now(),
		EmailRef:  "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	msgs := []events.PhishingEvent{expMsg}
	expUPE := awards.UserPhishingEvent{
		UserID:      awards.UserId(expMsg.UserId),
		Action:      awards.Opened,
		CreatedAt:   expMsg.CreatedAt,
		EmailRef:    expMsg.EmailRef,
		ProcessedAt: nil,
	}

	mockQueue.EXPECT().GetNextMessages().Return(msgs, nil)
	mockRepo.EXPECT().Write(expUPE).Return(nil)
	mockQueue.EXPECT().DeleteMessage(expMsg.EventId).Return(nil)

	mockQueue.EXPECT().GetNextMessages().Return(nil, nil)

	eL := eventlogging.New(mockRepo, mockQueue)
	eL.LogNewEvents()
}
