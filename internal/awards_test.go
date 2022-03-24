package awards_test

import (
	"fmt"
	awards "gosanta/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		Name          string
		User          awards.User
		emailRef      string
		action        awards.PhishingAction
		expectedErr   error
		expectedAward *awards.PhishingAward
	}{
		{
			Name:        "first award",
			User:        awards.User{Id: awards.UserId(1), CompanyId: awards.CompanyId(1)},
			emailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
			action:      awards.Opened,
			expectedErr: nil,
			expectedAward: &awards.PhishingAward{
				Id:         0,
				AssignedTo: awards.UserId(1),
				Type:       awards.OpenAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
		},
		{
			Name: "duplicate",
			User: awards.User{
				Id:        awards.UserId(1),
				CompanyId: awards.CompanyId(1),
				Awards: []awards.PhishingAward{
					{
						Id:         0,
						AssignedTo: awards.UserId(1),
						Type:       awards.OpenAward,
						EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
					},
				},
			},
			emailRef: "f20416ef-15d5-4159-9bef-de150edfa970",
			action:   awards.Opened,
			expectedErr: &awards.Error{
				Code: awards.DuplicateError,
				Err: fmt.Errorf(
					"user %v has already earned award from email %v",
					awards.UserId(1),
					"f20416ef-15d5-4159-9bef-de150edfa970",
				),
			},
			expectedAward: nil,
		},
		{
			Name: "new award",
			User: awards.User{
				Id:        awards.UserId(1),
				CompanyId: awards.CompanyId(1),
				Awards: []awards.PhishingAward{
					{
						Id:         0,
						AssignedTo: awards.UserId(1),
						Type:       awards.OpenAward,
						EmailRef:   "36b9c31a-d090-4e54-8660-6c44e2947aa0",
					},
				},
			},
			emailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
			action:      awards.Opened,
			expectedErr: nil,
			expectedAward: &awards.PhishingAward{
				Id:         0,
				AssignedTo: awards.UserId(1),
				Type:       awards.OpenAward,
				EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			a, err := awards.New(test.User, test.emailRef, test.action)

			assert.Equal(t, test.expectedErr, err)
			if test.expectedAward != nil {
				assert.Equal(t, test.expectedAward.AssignedTo, a.AssignedTo)
				assert.Equal(t, test.expectedAward.EmailRef, a.EmailRef)
				assert.Equal(t, test.expectedAward.Type, a.Type)
			}
		})
	}
}
