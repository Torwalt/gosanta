package awarding

import (
	"gosanta/internal"
	"gosanta/internal/ports"
)

type AwardService struct {
	awardRepo ports.AwardRepository
	userRepo  ports.UserReadRepository
}

func NewAwardService(awardRepo ports.AwardRepository, userRepo ports.UserReadRepository) AwardService {
	return AwardService{awardRepo: awardRepo, userRepo: userRepo}
}

// Assign a PhishingAward to a User.
func (s *AwardService) AssignPhishingAward(cpa *ports.CreatePhishingAward) (*awards.PhishingAward, error) {
	u, err := s.userRepo.Get(cpa.Id)
	if err != nil {
		return nil, err
	}
	a, err := awards.New(*u, cpa.EmailRef, cpa.Reason)
	if err != nil {
		return nil, err
	}
	a, err = s.awardRepo.Add(a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
