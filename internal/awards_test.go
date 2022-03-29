package awards_test

import (
	"fmt"
	awards "gosanta/internal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsProtected(t *testing.T) {
	tests := []struct {
		Name           string
		EarnedOn       time.Time
		ExpIsProtected bool
	}{
		{
			Name:           "is not protected",
			EarnedOn:       time.Now().AddDate(0, 0, -4),
			ExpIsProtected: false,
		},
		{
			Name:           "is protected",
			EarnedOn:       time.Now().AddDate(0, 0, -5),
			ExpIsProtected: true,
		},
	}

	for _, test := range tests {
		pa := &awards.PhishingAward{
			Id:         int64(1),
			AssignedTo: awards.UserId(1),
			EarnedOn:   test.EarnedOn,
			Type:       awards.OpenAward,
			EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
		}
		isProtected := pa.IsProtected()

		assert.Equal(t, test.ExpIsProtected, isProtected)
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		Name          string
		User          awards.User
		Event         awards.UserPhishingEvent
		WasClicked    bool
		ExistingAward *awards.PhishingAward
		ExpectedError error
		ExpectedAward *awards.PhishingAward
	}{
		{
			Name: "new award created",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Opened,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked:    false,
			ExistingAward: nil,
			ExpectedError: nil,
			ExpectedAward: &awards.PhishingAward{
				Id:         int64(1),
				AssignedTo: awards.UserId(1),
				EarnedOn:   time.Now(),
				Type:       awards.OpenAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
		},
		{
			Name: "existing award updated",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Reported,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked: false,
			ExistingAward: &awards.PhishingAward{
				Id:         int64(1),
				AssignedTo: awards.UserId(1),
				EarnedOn:   time.Now().Add(time.Duration(-10)),
				Type:       awards.OpenAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
			ExpectedError: nil,
			ExpectedAward: &awards.PhishingAward{
				Id:         int64(1),
				AssignedTo: awards.UserId(1),
				EarnedOn:   time.Now(),
				Type:       awards.ReportAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
		},
		{
			Name: "action not eligible for award",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Clicked,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked:    false,
			ExistingAward: nil,
			ExpectedError: &awards.Error{Code: awards.NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award", awards.Clicked)},
			ExpectedAward: nil,
		},
		{
			Name: "clicked event exists",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Opened,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked:    true,
			ExistingAward: nil,
			ExpectedError: &awards.Error{Code: awards.NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: phishing link was clicked", awards.Opened)},
			ExpectedAward: nil,
		},
		{
			Name: "already ignored",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Opened,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked:    false,
			ExistingAward: &awards.PhishingAward{
				Id:         int64(1),
				AssignedTo: awards.UserId(1),
				EarnedOn:   time.Now().Add(time.Duration(-10)),
				Type:       awards.IgnoreAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
			ExpectedError: &awards.Error{Code: awards.NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: email was already ignored", awards.Opened)},
			ExpectedAward: nil,
		},
		{
			Name: "is duplicate",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Opened,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked:    false,
			ExistingAward: &awards.PhishingAward{
				Id:         int64(1),
				AssignedTo: awards.UserId(1),
				EarnedOn:   time.Now().Add(time.Duration(-10)),
				Type:       awards.OpenAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
			ExpectedError: &awards.Error{Code: awards.NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: award already earned", awards.Opened)},
			ExpectedAward: nil,
		},
		{
			Name: "not upgradeable",
			User: awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			Event: awards.UserPhishingEvent{
				ID:          int64(1),
				UserID:      awards.UserId(1),
				Action:      awards.Opened,
				CreatedAt:   time.Now(),
				EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
				ProcessedAt: nil,
			},
			WasClicked:    false,
			ExistingAward: &awards.PhishingAward{
				Id:         int64(1),
				AssignedTo: awards.UserId(1),
				EarnedOn:   time.Now().Add(time.Duration(-10)),
				Type:       awards.ReportAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
			ExpectedError: &awards.Error{Code: awards.NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: cannot upgrade award", awards.Opened)},
			ExpectedAward: nil,
		},
	}

	for _, test := range tests {
		award, err := awards.New(test.User, test.Event, test.WasClicked, test.ExistingAward)

		assert.Equal(t, test.ExpectedError, err)

		if test.ExpectedAward == nil {
			assert.Nil(t, award)
		}

		if test.ExpectedAward != nil {
			assert.Equal(t, test.ExpectedAward.AssignedTo, award.AssignedTo)
			assert.Equal(t, test.ExpectedAward.Type, award.Type)
			assert.Equal(t, test.ExpectedAward.EmailRef, award.EmailRef)
		}
	}
}
