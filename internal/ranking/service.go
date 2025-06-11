package ranking

import (
	awards "gosanta/internal"
	"gosanta/internal/ports"
)

type RankingService struct {
	awardRepo ports.AwardReadRepository
	userRepo  ports.UserReadRepository
}

func NewService(
	awardRepo ports.AwardReadRepository,
	userRepo ports.UserReadRepository,
) RankingService {
	return RankingService{awardRepo: awardRepo, userRepo: userRepo}
}

// Return a PhishingAward for an Id.
func (s *RankingService) Get(id int64) (*awards.PhishingAward, error) {
	award, err := s.awardRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return award, nil
}

// Return all PhishingAwards for a User.
func (s *RankingService) GetUserAwards(uId awards.UserId) ([]awards.PhishingAward, error) {
	awardS, err := s.awardRepo.GetUserAwards(uId)
	if err != nil {
		return awardS, err
	}
	return awardS, nil
}

// Calculate a ranked Leaderboard for the passed user's company.
func (s *RankingService) CalcLeaderboard(uId awards.UserId) (*awards.Leaderboard, error) {
	leaderboard := &awards.Leaderboard{}

	user, err := s.userRepo.Get(uId)
	if err != nil {
		return leaderboard, err
	}

	userS, err := s.userRepo.GetCompanyUsers(user.CompanyId)
	if err != nil {
		return leaderboard, err
	}

	members := []awards.LeaderboardMember{}
	for _, u := range userS {
		lm := awards.AsLeaderboardMember(&u)
		members = append(members, lm)
	}
	leaderboard = awards.MakeLeaderboard(&members)

	return leaderboard, nil
}
