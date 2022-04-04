package awards

import (
	"sort"
)

// Award summary for a user.
type AwardSummary struct {
	IgnoringAward int
	OpenAward     int
	ReportAward   int
}

// Create an award summary for PhishingAwards.
func MakeAwardSummary(aS []PhishingAward) AwardSummary {
	var ignoreCount int
	var openCount int
	var reportCount int

	for _, award := range aS {
		switch award.Type {
		case IgnoreAward:
			ignoreCount += 1
		case OpenAward:
			openCount += 1
		case ReportAward:
			reportCount += 1
		}
	}

	return AwardSummary{
		IgnoringAward: ignoreCount,
		OpenAward:     openCount,
		ReportAward:   reportCount,
	}
}

// Summed up points of earned awards.
func (as *AwardSummary) Sum() int {
	return IgnoreAward.Points()*as.IgnoringAward + OpenAward.Points()*as.OpenAward + ReportAward.Points()*as.ReportAward
}

// User representation in a Leaderboard.
type LeaderboardMember struct {
	UserId       UserId
	UserFullName string
	Score        int
	Summary      AwardSummary
	Rank         int
}

// Create an unranked LeaderboardMember representation of a User.
func AsLeaderboardMember(u *User) LeaderboardMember {
	summary := MakeAwardSummary(u.Awards)
	return LeaderboardMember{
		UserId:       u.Id,
		UserFullName: u.FullName(),
		Score:        summary.Sum(),
		Summary:      summary,
		Rank:         0,
	}
}

// Leaderboard of Users of one company in context of a requesting user.
type Leaderboard struct {
	RankedUsers []LeaderboardMember
}

// Create a dense Leaderboard ranking for LeaderboardMembers.
func MakeLeaderboard(m *[]LeaderboardMember) *Leaderboard {
	rankedMembers := DenseRank(m)
	return &Leaderboard{RankedUsers: *rankedMembers}
}

// Sort and assign ranks to the LeaderboardMembers. Sorting is done on the
// Score in descending order. Supported ranking strategies: Dense.
//
// A Dense ranking strategy will assign the same
// rank multiple times if the Scores are the same.
// E.g. ranks: 1,1,2,3,3,3,4 etc.
func DenseRank(m *[]LeaderboardMember) *[]LeaderboardMember {
	if len(*m) == 0 {
		return m
	}

	// Ideally, rank assignment would be done in the same iteration. Such
	// improvements can be done later.
	ref := *m
	sort.Slice(ref, func(i, j int) bool {
		return ref[i].Score > ref[j].Score
	})

	// init first rank
	rank := 1
	curScore := ref[0].Score
	ref[0].Rank = rank
	rest := ref[1:]

	for idx := range rest {
		if curScore == rest[idx].Score {
			rest[idx].Rank = rank
			continue
		}
		curScore = rest[idx].Score
		rank += 1
		rest[idx].Rank = rank
	}

	return m
}
