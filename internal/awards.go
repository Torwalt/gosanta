package awards

import (
	"fmt"
	"time"
)

// An award cannot be upgraded or removed, if the related test phishing email
// is interacted with by a user *after* AwardProtectedAfterDays.
const AwardProtectedAfterDays = 5

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

// String representation of an AwardType.
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

// A PhishingAward cannot be upgraded or removed, if it was assigned and stayed
// assigned to a user for AwardProtectedAfterDays.
func (a *PhishingAward) IsProtected() bool {
	now := time.Now()
	protectedAt := now.AddDate(0, 0, -AwardProtectedAfterDays)
	if a.EarnedOn.Before(protectedAt) {
		return true
	}

	return false
}

// Prepare a new award for a user. An award can be assigned for an interaction of a User with a
// training phishing email.
//
// Actions if an award already exists:
// 1. Upgrade an OpenAward to a ReportAward. UserPhishingEvent must be of PhishingAction Opened.
//    The existing award should not be protected.
// 2. Remove existing award. If PhishingAction is Clicked and award not protected.
// 3. Do nothing. If user has existing IgnoreAward or award is protected or award exists already.
//
// Actions if no award exists:
// 1. Assign AwardType based on PhishingAction.
// 2. Do nothing. If PhishingAction is Clicked.
func New(
	u User,
	event UserPhishingEvent,
	clickedExists bool,
	existingAward *PhishingAward,
) (*PhishingAward, error) {
	typ := event.Action.ToAwardType()
	if typ == -1 {
		return nil, &Error{Code: NoAward, Err: fmt.Errorf(
			"action %v is not eligible for award", event.Action)}
	}

	// in any case, no awards here
	if clickedExists {
		return nil, &Error{Code: NoAward, Err: fmt.Errorf(
			"action %v is not eligible for award: phishing link was clicked", event.Action)}
	}

	if existingAward != nil {
		if wasIgnored(*existingAward) == true {
			return nil, &Error{Code: NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: email was already ignored", event.Action)}
		}
		if isDuplicate(*existingAward, typ) == true {
			return nil, &Error{Code: NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: award already earned", event.Action)}
		}
		// e.g. if ReportedAward was earned and OpenedEvent comes in.
		if isUpgradable(*existingAward, event.Action) == false {
			return nil, &Error{Code: NoAward, Err: fmt.Errorf(
				"action %v is not eligible for award: cannot upgrade award", event.Action)}
		}
	}

	return &PhishingAward{
		AssignedTo: u.Id,
		EarnedOn:   time.Now(),
		Type:       typ,
		EmailRef:   event.EmailRef,
	}, nil
}

func isUpgradable(existingAward PhishingAward, action PhishingAction) bool {
	if existingAward.Type == OpenAward && existingAward.IsProtected() == false && action == Reported {
		return true
	}
	return false
}

func isDuplicate(award PhishingAward, awardType AwardType) bool {
	if award.Type == awardType {
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
