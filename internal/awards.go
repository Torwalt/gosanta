package awards

import (
	"fmt"
	"time"
)

// The reason an award was given.
type AwardType int

const (
	// For opening a phishing mail
	OpenAward AwardType = iota
	// For reporting a phishing mail
	ReportAward
	// For ignoring a phishing mail
	IgnoreAward
)

func (a AwardType) String() string {
	switch a {
	case OpenAward:
		return "opened award"
	case ReportAward:
		return "reported award"
	case IgnoreAward:
		return "ignored award"
	default:
		return ""
	}
}

// An award assigned for correctly interacting with a pyphish phishing email.
type PhishingAward struct {
	Id         int64
	AssignedTo UserId
	EarnedOn   time.Time
	Type       AwardType
	EmailRef   string
}

// Prepare a new award for a user
func New(u User, emailRef string, action PhishingAction) (*PhishingAward, error) {
	typ := action.ToAwardType()
	if typ == -1 {
		return nil, &Error{Code: NoAward, Err: fmt.Errorf(
			"action %v is not eligible for award", action.String())}
	}

	award := u.FindRelatedAward(emailRef)
	if award != nil {
		if err := canAssignAward(*award, typ); err != nil {
			return nil, err
		}
	}

	return &PhishingAward{
		AssignedTo: u.Id,
		EarnedOn:   time.Now(),
		Type:       typ,
		EmailRef:   emailRef,
	}, nil
}

func canAssignAward(existingAward PhishingAward, newAward AwardType) error {
	if isDuplicate(existingAward, newAward) {
		return &Error{Code: DuplicateError, Err: fmt.Errorf(
			"user %v has already earned award from email %v", existingAward.AssignedTo, existingAward.EmailRef)}
	}
	if wasIgnored(existingAward) {
		return &Error{Code: NoAward, Err: fmt.Errorf(
			"user %v has ignored email %v for too long", existingAward.AssignedTo, existingAward.EmailRef)}
	}
	// if wasClicked(existingAward) {
	// 	return &Error{Code: NoAward, Err: fmt.Errorf(
	// 		"user %v has already clicked the phishing link in email %v", existingAward.AssignedTo, existingAward.EmailRef)}
	// }
	return nil
}

func isDuplicate(award PhishingAward, newAward AwardType) bool {
	if award.Type == newAward {
		return true
	}
	return false
}

func wasIgnored(award PhishingAward) bool {
	if award.Type == IgnoreAward {
		return true
	}
	return false
}

// func wasClicked(award PhishingAward) bool {
// 	if award.Type == Clicked {
// 		return true
// 	}
// 	return false
// }
