package rest_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/mocks"
	"gosanta/internal/rest"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserAwards(t *testing.T) {
	userId := awards.UserId(1)
	url := "/awards/v1/user/1"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	ctrl := gomock.NewController(t)
	arSrv := mocks.NewMockAwardReadingService(ctrl)

	awardS := []awards.PhishingAward{
		{
			Id:         1,
			AssignedTo: userId,
			EarnedOn:   time.Now(),
			Type:       awards.OpenAward,
			EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
		},
	}
	arSrv.EXPECT().GetUserAwards(userId).Return(awardS, nil)
	srv := rest.New(arSrv)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	gotS := []rest.UserAwardResponse{}
	want := rest.UserAwardResponse{
		Id:        awardS[0].Id,
		UserId:    int(awardS[0].AssignedTo),
		CreatedOn: awardS[0].EarnedOn,
		Type:      awardS[0].Type.String(),
	}
	err = json.NewDecoder(rr.Body).Decode(&gotS)
	if err != nil {
		t.Errorf("could not decode response body: %v", err)
	}
	got := gotS[0]

	assert.Equal(t, want.Id, got.Id)
	assert.Equal(t, want.UserId, got.UserId)
	assert.Equal(t, want.Type, got.Type)
	assert.WithinDuration(t, want.CreatedOn, got.CreatedOn, 0)
}

func TestGetUserAwardsErrors(t *testing.T) {
	tests := []struct {
		Name         string
		UserId       string
		StatusCode   int
		ActualUserId awards.UserId
		ServiceCall  bool
		ServiceError error
	}{
		{
			Name:         "malformed user id",
			UserId:       "asd",
			StatusCode:   http.StatusBadRequest,
			ActualUserId: awards.UserId(1),
			ServiceCall:  false,
		},
		{
			Name:         "not found",
			UserId:       "1",
			StatusCode:   http.StatusNotFound,
			ActualUserId: awards.UserId(1),
			ServiceCall:  true,
			ServiceError: &awards.Error{Code: awards.DoesNotExistError, Err: fmt.Errorf("Test error")},
		},
		{
			Name:         "unexpected error",
			UserId:       "1",
			StatusCode:   http.StatusInternalServerError,
			ActualUserId: awards.UserId(1),
			ServiceCall:  true,
			ServiceError: fmt.Errorf("Test error"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := "/awards/v1/user/" + test.UserId
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Errorf("Error creating a new request: %v", err)
			}

			rr := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			arSrv := mocks.NewMockAwardReadingService(ctrl)

			awardS := []awards.PhishingAward{
				{
					Id:         1,
					AssignedTo: test.ActualUserId,
					EarnedOn:   time.Now(),
					Type:       awards.OpenAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
				},
			}

			if test.ServiceCall {
				arSrv.EXPECT().GetUserAwards(test.ActualUserId).Return(awardS, test.ServiceError)
			}

			srv := rest.New(arSrv)
			srv.ServeHTTP(rr, req)
			assert.Equal(t, test.StatusCode, rr.Code)
		})
	}
}

func TestCalcLeaderboard(t *testing.T) {
	userId := awards.UserId(1)
	url := "/awards/v1/user/1/leaderboard"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	ctrl := gomock.NewController(t)
	arSrv := mocks.NewMockAwardReadingService(ctrl)

	lb := &awards.Leaderboard{
		RankedUsers: []awards.LeaderboardMember{
			{
				UserId:       userId,
				UserFullName: "Roboute Guilliman",
				Score:        60,
				Summary: awards.AwardSummary{
					IgnoringAward: 0,
					OpenAward:     0,
					ReportAward:   10,
				},
				Rank: 1,
			},
		},
	}
	arSrv.EXPECT().CalcLeaderboard(userId).Return(lb, nil)
	srv := rest.New(arSrv)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	gotS := []rest.LeaderboardMember{}
	wantS := []rest.LeaderboardMember{
		{
			UserId:       int(userId),
			UserFullName: "Roboute Guilliman",
			Score:        60,
			IgnoreCount:  0,
			OpenCount:    0,
			ReportCount:  10,
			Rank:         1,
		},
	}
	err = json.NewDecoder(rr.Body).Decode(&gotS)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, wantS, gotS)
}
