package awards_test

import (
	"testing"
	"time"

	awards "gosanta/internal"

	"github.com/stretchr/testify/assert"
)

func TestMakeAwardSummary(t *testing.T) {
	tests := []struct {
		Name            string
		AwardS          []awards.PhishingAward
		ExpectedSummary awards.AwardSummary
	}{
		{
			Name: "2 ignores 1 open 1 reported",
			AwardS: []awards.PhishingAward{
				{
					Id:         int64(1),
					AssignedTo: awards.UserId(1),
					EarnedOn:   time.Now(),
					Type:       awards.IgnoreAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
				},
				{
					Id:         int64(1),
					AssignedTo: awards.UserId(1),
					EarnedOn:   time.Now().AddDate(0, 0, -5),
					Type:       awards.IgnoreAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa971",
				},
				{
					Id:         int64(1),
					AssignedTo: awards.UserId(1),
					EarnedOn:   time.Now().AddDate(0, 0, -3),
					Type:       awards.OpenAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa391",
				},
				{
					Id:         int64(1),
					AssignedTo: awards.UserId(1),
					EarnedOn:   time.Now().AddDate(0, 0, -2),
					Type:       awards.ReportAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa413",
				},
			},
			ExpectedSummary: awards.AwardSummary{
				IgnoringAward: 2,
				OpenAward:     1,
				ReportAward:   1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			as := awards.MakeAwardSummary(test.AwardS)
			assert.Equal(t, test.ExpectedSummary, as)
		})
	}
}

func TestAwardSummarySum(t *testing.T) {
	as := awards.AwardSummary{
		IgnoringAward: 2,
		OpenAward:     3,
		ReportAward:   5,
	}
	actual := as.Sum()
	expected := (1 * 2) + (2 * 3) + (3 * 5)

	assert.Equal(t, expected, actual)
}

func TestLeaderboardDenseRank(t *testing.T) {
	tests := []struct {
		Name      string
		ExpRanks  []int
		ExpScores []int
		Lmembers  []awards.LeaderboardMember
	}{
		{
			Name:      "1, 2, 2, 3",
			ExpRanks:  []int{1, 2, 2, 3},
			ExpScores: []int{20, 15, 15, 10},
			Lmembers: []awards.LeaderboardMember{
				{
					UserId:       awards.UserId(1),
					UserFullName: "Rank 3",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(2),
					UserFullName: "Rank 2",
					Score:        15,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     1,
						ReportAward:   1,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(3),
					UserFullName: "Rank 2",
					Score:        15,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     1,
						ReportAward:   1,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(4),
					UserFullName: "Rank 1",
					Score:        20,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     2,
						ReportAward:   2,
					},
					Rank: 0,
				},
			},
		},
		{
			Name:      "1, 1, 1, 1",
			ExpRanks:  []int{1, 1, 1, 1},
			ExpScores: []int{10, 10, 10, 10},
			Lmembers: []awards.LeaderboardMember{
				{
					UserId:       awards.UserId(1),
					UserFullName: "Rank 1",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(2),
					UserFullName: "Rank 1",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(3),
					UserFullName: "Rank 1",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(4),
					UserFullName: "Rank 1",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
			},
		},
		{
			Name:      "1, 2, 2, 2",
			ExpRanks:  []int{1, 2, 2, 2},
			ExpScores: []int{20, 10, 10, 10},
			Lmembers: []awards.LeaderboardMember{
				{
					UserId:       awards.UserId(1),
					UserFullName: "Rank 1",
					Score:        20,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     2,
						ReportAward:   2,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(2),
					UserFullName: "Rank 2",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(3),
					UserFullName: "Rank 2",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
				{
					UserId:       awards.UserId(4),
					UserFullName: "Rank 3",
					Score:        10,
					Summary: awards.AwardSummary{
						IgnoringAward: 10,
						OpenAward:     0,
						ReportAward:   0,
					},
					Rank: 0,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rankedMembers := awards.DenseRank(&test.Lmembers)
			expMembers := len(test.ExpRanks)
			rm := *rankedMembers

			assert.Equal(t, expMembers, len(rm))
			for n := 0; n < expMembers; n++ {
				assert.Equal(t, rm[n].Rank, test.ExpRanks[n])
				assert.Equal(t, rm[n].Score, test.ExpScores[n])
			}
		})
	}
}
