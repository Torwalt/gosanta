package awarding

import (
	awards "gosanta/internal"
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

// Assign or remove awards for the corresponding users based on the UserPhishingEvent.
func (s *AwardService) AssignAward(event awards.UserPhishingEvent) (usrAwardEvent awards.UserAwardEvent, err error) {
	usrAwardEvent, err = s.assignPhishingAward(event)
	if err != nil {
		return usrAwardEvent, err
	}
	return usrAwardEvent, nil
}

func (s *AwardService) assignPhishingAward(event awards.UserPhishingEvent) (awards.UserAwardEvent, error) {
	userAward := awards.UserAwardEvent{Event: event}
	clickedExists, err := s.eventRepo.ClickedExists(event.UserID, event.EmailRef)
	if err != nil {
		// TODO
		return userAward, err
	}

	existingAward, err := s.awardRepo.GetByEmailRef(event.UserID, event.EmailRef)
	if err != nil {
		// TODO
		return userAward, err
	}

	user, err := s.userRepo.Get(event.UserID)
	if err != nil {
		// TODO
		return userAward, err
	}

	// Remove existing award if clicked and not protected.
	if event.Action == awards.Clicked {
		if existingAward != nil && !existingAward.IsProtected() {
			err := s.awardRepo.Delete(existingAward.Id)
			if err != nil {
				return userAward, err
			}
		}
		//
		return userAward, nil
	}

	newAward, err := awards.New(*user, event, clickedExists, existingAward)
	if err != nil {
		// TODO
		return userAward, err
	}

	userAward.Award = newAward

	// Award has been upgraded, no new award is created.
	if existingAward != nil && newAward != nil {
		err := s.awardRepo.UpdateExisting(existingAward, newAward)
		return userAward, err
	}

	err = s.awardRepo.Add(newAward)
	return userAward, err
}
