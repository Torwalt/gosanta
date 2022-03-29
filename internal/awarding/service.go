package awarding

import (
	"fmt"
	"gosanta/internal"
	"gosanta/internal/ports"
)

type AwardService struct {
	awardRepo ports.AwardRepository
	userRepo  ports.UserReadRepository
	eventRepo ports.EventRepositoryProcessor
}

func NewAwardService(
	awardRepo ports.AwardRepository,
	userRepo ports.UserReadRepository,
	eventRepo ports.EventRepositoryProcessor,
) AwardService {
	return AwardService{awardRepo: awardRepo, userRepo: userRepo, eventRepo: eventRepo}
}

// Retrieve unprocessed PhishingEvents and assign or remove awards for the
// corresponding users based on the interaction with a test phishing mail.
func (s *AwardService) ProcessPhishingEvents() error {
	events, err := s.eventRepo.GetUnprocessed()
	if err != nil {
		return fmt.Errorf("could not retrieve unprocessed events: %v", err)
	}

	for _, newEvent := range events {
		err = s.assignPhishingAward(newEvent)
		if err != nil {
			// TODO
			continue
		}
		err := s.eventRepo.MarkAsProcessed(&newEvent)
		if err != nil {
			// TODO
			continue
		}
	}
	return nil
}

func (s *AwardService) assignPhishingAward(event awards.UserPhishingEvent) error {
	clickedExists, err := s.eventRepo.ClickedExists(event.UserID, event.EmailRef)
	if err != nil {
		// TODO
		return err
	}

	existingAward, err := s.awardRepo.GetByEmailRef(event.UserID, event.EmailRef)
	if err != nil {
		// TODO
		return err
	}

	user, err := s.userRepo.Get(event.UserID)
	if err != nil {
		// TODO
		return err
	}

	// Remove existing award if clicked and not protected.
	if event.Action == awards.Clicked {
		if existingAward != nil && !existingAward.IsProtected() {
			err := s.awardRepo.Delete(existingAward.Id)
			if err != nil {
				return err
			}
		}
		return nil
	}

	newAward, err := awards.New(*user, event, clickedExists, existingAward)
	if err != nil {
		// TODO
		return err
	}

	// Award has been upgraded, no new award is created.
	if existingAward != nil && newAward != nil {
		err := s.awardRepo.UpdateExisting(existingAward, newAward)
		return err
	}

	err = s.awardRepo.Add(newAward)
	return err
}
