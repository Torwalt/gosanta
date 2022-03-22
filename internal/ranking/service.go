package ranking

import (
	awards "gosanta/internal"
	"gosanta/internal/ports"
)

type RankingService struct {
	awardRepo ports.AwardReadRepository
	userRepo  ports.UserReadRepository
}

func New(awardRepo ports.AwardReadRepository, userRepo ports.UserReadRepository) RankingService {
	return RankingService{awardRepo: awardRepo, userRepo: userRepo}
}

func (s *RankingService) Get(u string) (awards.PhishingAward, error) {
	return awards.PhishingAward{}, nil
}

func (s *RankingService) GetUserAwards(uId awards.UserId) ([]awards.PhishingAward, error) {
	return []awards.PhishingAward{}, nil
}

func (s *RankingService) GetCompanyAwards(cId awards.CompanyId) ([]awards.PhishingAward, error) {
	return []awards.PhishingAward{}, nil
}
