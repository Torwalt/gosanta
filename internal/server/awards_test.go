package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/mocks"
	"gosanta/internal/server"

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
			Type:     awards.OpenAward,
			EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
		},
	}
	arSrv.EXPECT().GetUserAwards(userId).Return(awardS, nil)
	srv := server.New(arSrv)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	gotS := []server.UserAwardResponse{}
	want := server.UserAwardResponse{
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
					Type:     awards.OpenAward,
					EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
				},
			}

			if test.ServiceCall {
				arSrv.EXPECT().GetUserAwards(test.ActualUserId).Return(awardS, test.ServiceError)
			}

			srv := server.New(arSrv)
			srv.ServeHTTP(rr, req)
			assert.Equal(t, test.StatusCode, rr.Code)
		})
	}

}
