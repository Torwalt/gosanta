package eventbroker

import (
	"fmt"
	awards "gosanta/internal"
	"gosanta/internal/ports"
	"time"
)

// potentially do event persistence here?
type AwarderNotifier struct {
	eventLog       ports.EventLogReader
	awardService   ports.AwardAssigner
	mailSender     ports.MailSender
	eventPublisher ports.EventPublisher
	LogEvents      bool
}

func NewAwarderNotifier(eventLog ports.EventLogReader, awardService ports.AwardAssigner, mailSender ports.MailSender, eventPublisher ports.EventPublisher) AwarderNotifier {
	return AwarderNotifier{
		eventLog:       eventLog,
		awardService:   awardService,
		mailSender:     mailSender,
		eventPublisher: eventPublisher,
		LogEvents:      true,
	}
}

// Start the AwardNotifier.
// Concurrently run event retrieval/persitence, award assignment and user/system notification.
func (a *AwarderNotifier) Start(eventChan chan awards.UserPhishingEvent, awardChan chan awards.UserAwardEvent) {
	// poll and persist events
	go a.startEventLogging(eventChan)

	// process event and apply award assignment logic
	go a.startAwardAssigning(eventChan, awardChan)

	// apply logic in sending email to user and publish award event
	go a.startNotifying(awardChan)
}

func (a *AwarderNotifier) startEventLogging(eventChan chan awards.UserPhishingEvent) {
	for a.LogEvents {
		// TODO real logging
		fmt.Println("LogNewEvents called")
		events, err := a.eventLog.LogNewEvents()
		if err != nil {
			fmt.Printf("error when retrieving events: %v", err)
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
		fmt.Println("AssignAward called")
		awardEvnt, err := a.awardService.AssignAward(event)
		if err != nil {
			continue
		}
		outChan <- awardEvnt
	}
}

func (a *AwarderNotifier) startNotifying(eventChan chan awards.UserAwardEvent) {
	// SendToUser and PublishEvent can also be made concurrent, for now synchronious
	for awardEvnt := range eventChan {
		fmt.Println("SendToUser called")
		err := a.mailSender.SendToUser(awardEvnt)
		if err != nil {
			// TODO
			continue
		}
		fmt.Println("PublishEvent called")
		err = a.eventPublisher.PublishEvent(awardEvnt)
		if err != nil {
			// TODO
			continue
		}
	}
}