package eventbroker

import (
	"fmt"
	awards "gosanta/internal"
	"gosanta/internal/ports"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

// potentially do event persistence here?
type AwarderNotifier struct {
	eventLog       ports.EventLogReader
	awardService   ports.AwardAssigner
	mailSender     ports.MailSender
	eventPublisher ports.EventPublisher
	logger         log.Logger
	eventChan      chan awards.UserPhishingEvent
	awardChan      chan awards.UserAwardEvent
	LogEvents      bool
}

func NewAwarderNotifier(eventLog ports.EventLogReader, awardService ports.AwardAssigner,
	mailSender ports.MailSender, eventPublisher ports.EventPublisher, logger log.Logger,
	eventChan chan awards.UserPhishingEvent, awardChan chan awards.UserAwardEvent,
) AwarderNotifier {
	return AwarderNotifier{
		eventLog:       eventLog,
		awardService:   awardService,
		mailSender:     mailSender,
		eventPublisher: eventPublisher,
		LogEvents:      true,
		logger:         logger,
		eventChan:      eventChan,
		awardChan:      awardChan,
	}
}

// Start the AwardNotifier.
// Concurrently run event retrieval/persitence, award assignment and user/system notification.
func (a *AwarderNotifier) Start() {
	// poll and persist events
	go a.startEventLogging(a.eventChan)

	// process event and apply award assignment logic
	go a.startAwardAssigning(a.eventChan, a.awardChan)

	// apply logic in sending email to user and publish award event
	go a.startNotifying(a.awardChan)
}

func (a *AwarderNotifier) startEventLogging(eventChan chan awards.UserPhishingEvent) {
	// we should additionally retrieve events that were not processed due to process dying
	for a.LogEvents {
		events, err := a.eventLog.LogNewEvents()
		if err != nil {
			level.Error(a.logger).Log("error", err)
			time.Sleep(30 * time.Second)
			continue
		}
		for _, e := range events {
			eventChan <- e
		}
	}
}

func (a *AwarderNotifier) startAwardAssigning(inChan chan awards.UserPhishingEvent, outChan chan awards.UserAwardEvent) {
	for event := range inChan {
		awardEvnt, err := a.awardService.AssignAward(event)
		if err != nil {
			level.Error(a.logger).Log("error", err, "usecase", "awarding")
			a.scheduleAwardRetry(time.Now().UTC().Add(time.Second*60), event)
			continue
		}
		outChan <- awardEvnt
	}
}

func (a *AwarderNotifier) startNotifying(eventChan chan awards.UserAwardEvent) {
	// SendToUser and PublishEvent can also be made concurrent, for now synchronious
	for awardEvnt := range eventChan {
		// retrying would be implemented by the infrastructure component of the usecase
		// business logic would dictate, if event should be published first, or email to user
		err := a.mailSender.SendToUser(awardEvnt)
		if err != nil {
			level.Error(a.logger).Log("error", err)
			continue
		}
		err = a.eventPublisher.PublishEvent(awardEvnt)
		if err != nil {
			level.Error(a.logger).Log("error", err)
			continue
		}
	}
}

// downside: when process dies then the retry is lost
func (a *AwarderNotifier) scheduleAwardRetry(retryAt time.Time, event awards.UserPhishingEvent) {
	now := time.Now().UTC()
	waitTime := retryAt.Sub(now)

	msg := fmt.Sprintf("retrying in %v seconds", waitTime)
	level.Info(a.logger).Log("usecase", "awarding", "message", msg)

	f := func() {
		a.eventChan <- event
	}

	go a.retryIn(waitTime, f)
}

func (a *AwarderNotifier) retryIn(s time.Duration, f func()) {
	time.Sleep(s)
	f()
}
