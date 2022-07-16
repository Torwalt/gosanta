package eventbroker_test

import (
	awards "gosanta/internal"
	"gosanta/internal/eventbroker"
	"gosanta/internal/mocks"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go-kit/log"
	"github.com/golang/mock/gomock"
)

func TestStart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	el := mocks.NewMockEventLogReader(ctrl)
	as := mocks.NewMockAwardAssigner(ctrl)
	ms := mocks.NewMockMailSender(ctrl)
	ep := mocks.NewMockEventPublisher(ctrl)
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	eventChan := make(chan awards.UserPhishingEvent)
	awardChan := make(chan awards.UserAwardEvent)

	an := eventbroker.NewAwarderNotifier(el, as, ms, ep,
		log.With(logger, "component", "test-award-notifier"), eventChan, awardChan)

	uID := awards.UserId(1)
	expEvent := awards.UserPhishingEvent{
		ID:          1,
		UserID:      uID,
		Action:      awards.Opened,
		CreatedAt:   time.Now(),
		EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
		ProcessedAt: nil,
	}
	expEvents := []awards.UserPhishingEvent{}
	eventCount := 10
	for i := 0; i < eventCount; i++ {
		expEvents = append(expEvents, expEvent)
	}

	wg := sync.WaitGroup{}
	// LogNewEvents is called once, AssignAward, SendToUser and PublishEvent
	// are called eventCount times.
	wgCount := 1 + (eventCount * 3)
	wg.Add(wgCount)

	el.EXPECT().LogNewEvents().Return(expEvents, nil).Do(func() {
		defer wg.Done()
		an.LogEvents = false
	}).MaxTimes(1)

	expAward := awards.PhishingAward{
		Id:         1,
		AssignedTo: uID,
		EarnedOn:   time.Now(),
		Type:       awards.OpenAward,
		EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	expAwardEvent := awards.UserAwardEvent{
		Event: expEvent,
		Award: &expAward,
	}

	as.EXPECT().AssignAward(expEvents[0]).Return(expAwardEvent, nil).Do(func(arg0 interface{}) {
		defer wg.Done()
	}).AnyTimes()
	ms.EXPECT().SendToUser(expAwardEvent).Do(func(arg0 interface{}) {
		defer wg.Done()
	}).AnyTimes()
	ep.EXPECT().PublishEvent(expAwardEvent).Do(func(arg0 interface{}) {
		defer wg.Done()
	}).AnyTimes()

	an.Start()
	wg.Wait()
}
