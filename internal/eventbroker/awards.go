package eventbroker

import (
	"fmt"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/ports"

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
func (a *AwarderNotifier) Start() error {
	// process event and apply award assignment logic
	go a.startAwardAssigning(a.eventChan, a.awardChan)

	// apply logic in sending email to user and publish award event
	go a.startNotifying(a.awardChan)

	// On startup, first process unprocessed events once. That guarantees that
	// scheduled retries lost due to program shutdown are still processed.
	// Scheduled time is not considered.
	err := a.processUnprocessedEvents(a.eventChan)
	if err != nil {
		return fmt.Errorf("could not process unprocessed events: %v", err)
	}

	// poll and persist events
	go a.startEventLogging(a.eventChan)

	return nil
}

func (a *AwarderNotifier) processUnprocessedEvents(eventChan chan awards.UserPhishingEvent) error {
	events, err := a.eventLog.GetUnprocessedEvents()
	if err != nil {
		return err
	}

	for _, event := range events {
		eventChan <- event
	}
	return nil
}

func (a *AwarderNotifier) startEventLogging(eventChan chan awards.UserPhishingEvent) {
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

// Schedule a retry for processing a UserPhishingEvent. Retries are kept in
// memory and are lost on program shutdown.
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
