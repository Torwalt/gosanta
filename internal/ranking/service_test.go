package ranking_test

import (
	"testing"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/mocks"
	"gosanta/internal/ranking"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessPhishingEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	ar := mocks.NewMockAwardReadRepository(ctrl)
	ur := mocks.NewMockUserReadRepository(ctrl)
	rs := ranking.NewService(ar, ur)

	reqUID := awards.UserId(1)
	compID := awards.CompanyId(1)
	reqUser := &awards.User{
		Id:        reqUID,
		FirstName: "Bob",
		LastName:  "Bobacious",
		CompanyId: compID,
		Awards: []awards.PhishingAward{
			{
				Id:         int64(1),
				AssignedTo: reqUID,
				EarnedOn:   time.Now(),
				Type:       awards.OpenAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
		},
	}
	ur.EXPECT().Get(reqUID).Return(reqUser, nil)

	expCompUsers := []awards.User{
		*reqUser,
		{
			Id:        awards.UserId(2),
			FirstName: "Fred",
			LastName:  "Fredacious",
			CompanyId: compID,
			Awards: []awards.PhishingAward{
				{
					Id:         int64(2),
					AssignedTo: awards.UserId(2),
					EarnedOn:   time.Now(),
					Type:       awards.OpenAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
				},
			},
		},
		{
			Id:        awards.UserId(3),
			FirstName: "Rob",
			LastName:  "Robacious",
			CompanyId: compID,
			Awards: []awards.PhishingAward{
				{
					Id:         int64(3),
					AssignedTo: awards.UserId(3),
					EarnedOn:   time.Now(),
					Type:       awards.ReportAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
				},
			},
		},
		{
			Id:        awards.UserId(4),
			FirstName: "Jack",
			LastName:  "Jackacious",
			CompanyId: compID,
			Awards: []awards.PhishingAward{
				{
					Id:         int64(4),
					AssignedTo: awards.UserId(4),
					EarnedOn:   time.Now(),
					Type:       awards.IgnoreAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
				},
			},
		},
	}
	ur.EXPECT().GetCompanyUsers(compID).Return(expCompUsers, nil)

	lb, err := rs.CalcLeaderboard(reqUID)

	assert.Nil(t, err)
	assert.NotNil(t, lb)

	assert.Equal(t, 4, len(lb.RankedUsers))
	assert.Equal(t, 1, lb.RankedUsers[0].Rank)
	assert.Equal(t, 2, lb.RankedUsers[1].Rank)
	assert.Equal(t, 2, lb.RankedUsers[2].Rank)
	assert.Equal(t, 3, lb.RankedUsers[3].Rank)
}
