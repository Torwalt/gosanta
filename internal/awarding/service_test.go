package awarding_test

import (
	awards "gosanta/internal"
	"gosanta/internal/awarding"
	"gosanta/internal/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessPhishingEventsAwardUpdated(t *testing.T) {
	ctrl := gomock.NewController(t)
	ar := mocks.NewMockAwardRepository(ctrl)
	ur := mocks.NewMockUserReadRepository(ctrl)
	er := mocks.NewMockEventRepositoryProcessor(ctrl)
	user_1 := awards.UserId(1)

	event := awards.UserPhishingEvent{
		ID:          1,
		UserID:      user_1,
		Action:      awards.Reported,
		CreatedAt:   time.Now().Add(time.Duration(-100)),
		EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
		ProcessedAt: nil,
	}
	unproEvents := []awards.UserPhishingEvent{event}
	er.EXPECT().GetUnprocessed().Times(1).Return(unproEvents, nil)
	er.EXPECT().ClickedExists(event.UserID, event.EmailRef).Return(false, nil)

	existAward := &awards.PhishingAward{
		Id:         int64(1),
		AssignedTo: user_1,
		EarnedOn:   time.Now().AddDate(0, 0, -1),
		Type:       awards.OpenAward,
		EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	ar.EXPECT().GetByEmailRef(event.UserID, event.EmailRef).Return(existAward, nil)

	user := &awards.User{
		Id:        user_1,
		CompanyId: awards.CompanyId(1),
	}
	ur.EXPECT().Get(event.UserID).Return(user, nil)

	ar.EXPECT().UpdateExisting(existAward, gomock.Any()).Return(nil)
	ar.EXPECT().Add(gomock.Any()).Times(0)

	er.EXPECT().MarkAsProcessed(gomock.Any()).Return(nil)

	as := awarding.NewAwardService(ar, ur, er)
	err := as.ProcessPhishingEvents()

	assert.Nil(t, err)
}

func TestProcessPhishingEventsAwardRemoveExisting(t *testing.T) {
	ctrl := gomock.NewController(t)
	ar := mocks.NewMockAwardRepository(ctrl)
	ur := mocks.NewMockUserReadRepository(ctrl)
	er := mocks.NewMockEventRepositoryProcessor(ctrl)
	user_1 := awards.UserId(1)

	event := awards.UserPhishingEvent{
		ID:          1,
		UserID:      user_1,
		Action:      awards.Clicked,
		CreatedAt:   time.Now().Add(time.Duration(-100)),
		EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
		ProcessedAt: nil,
	}
	unproEvents := []awards.UserPhishingEvent{event}
	er.EXPECT().GetUnprocessed().Times(1).Return(unproEvents, nil)
	er.EXPECT().ClickedExists(event.UserID, event.EmailRef).Return(false, nil)

	existAward := &awards.PhishingAward{
		Id:         int64(1),
		AssignedTo: user_1,
		// I.e. not protected
		EarnedOn: time.Now().AddDate(0, 0, -3),
		Type:     awards.OpenAward,
		EmailRef: "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	ar.EXPECT().GetByEmailRef(event.UserID, event.EmailRef).Return(existAward, nil)

	user := &awards.User{
		Id:        user_1,
		CompanyId: awards.CompanyId(1),
	}
	ur.EXPECT().Get(event.UserID).Return(user, nil)
	ar.EXPECT().Delete(existAward.Id).Return(nil)
	er.EXPECT().MarkAsProcessed(gomock.Any()).Return(nil)

	as := awarding.NewAwardService(ar, ur, er)
	err := as.ProcessPhishingEvents()

	assert.Nil(t, err)
}

func TestProcessPhishingEventsAwardAddNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	ar := mocks.NewMockAwardRepository(ctrl)
	ur := mocks.NewMockUserReadRepository(ctrl)
	er := mocks.NewMockEventRepositoryProcessor(ctrl)
	user_1 := awards.UserId(1)

	event := awards.UserPhishingEvent{
		ID:          1,
		UserID:      user_1,
		Action:      awards.Opened,
		CreatedAt:   time.Now().Add(time.Duration(-100)),
		EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
		ProcessedAt: nil,
	}
	unproEvents := []awards.UserPhishingEvent{event}
	er.EXPECT().GetUnprocessed().Times(1).Return(unproEvents, nil)
	er.EXPECT().ClickedExists(event.UserID, event.EmailRef).Return(false, nil)
	ar.EXPECT().GetByEmailRef(event.UserID, event.EmailRef).Return(nil, nil)

	user := &awards.User{
		Id:        user_1,
		CompanyId: awards.CompanyId(1),
	}
	ur.EXPECT().Get(event.UserID).Return(user, nil)
	er.EXPECT().MarkAsProcessed(gomock.Any()).Return(nil)
	ar.EXPECT().Add(gomock.Any()).Return(nil)

	as := awarding.NewAwardService(ar, ur, er)
	err := as.ProcessPhishingEvents()

	assert.Nil(t, err)
}
