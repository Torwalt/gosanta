package awards

import (
	"fmt"
	"time"
)

type UserId int64
type CompanyId int64

type PhishingAction int

const (
	Opened = iota
	Clicked
	Reported
	Ignored
)

func (p PhishingAction) ToAwardType() AwardType {
	switch p {
	case Opened:
		return OpenAward
	case Reported:
		return ReportAward
	case Ignored:
		return IgnoreAward
	default:
		return -1
	}
}

func (a PhishingAction) String() string {
	switch a {
	case Opened:
		return "opened"
	case Reported:
		return "reported"
	case Ignored:
		return "ignored"
	case Clicked:
		return "clicked"
	default:
		return ""
	}
}

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

// A user that can receive an award.
type User struct {
	Id        UserId
	CompanyId CompanyId
	Awards    []PhishingAward
}

// Return award received for the interaction with a phishing mail.
func (u *User) FindRelatedAward(emailRef string) *PhishingAward {
	for _, a := range u.Awards {
		if a.EmailRef == emailRef {
			return &a
		}
	}
	return nil
}

// A company to which users belong.
type Company struct {
	id    CompanyId
	users []User
}

// Prepare a new award for a user. An award cannot be gained twice.
func New(u User, emailRef string, action PhishingAction) (*PhishingAward, error) {
	typ := action.ToAwardType()
	if typ == -1 {
		return nil, &Error{Code: NoAward, Err: fmt.Errorf(
			"action %v is not eligible for award", action.String())}
	}

	award := u.FindRelatedAward(emailRef)
	if isDuplicate(award, typ) {
		return nil, &Error{Code: DuplicateError, Err: fmt.Errorf(
			"user %v has already earned award from email %v", u.Id, emailRef)}
	}

	return &PhishingAward{
		AssignedTo: u.Id,
		EarnedOn:   time.Now(),
		Type:       typ,
		EmailRef:   emailRef,
	}, nil
}

func isDuplicate(award *PhishingAward, newAward AwardType) bool {
	if award == nil {
		return false
	}
	if award.Type == newAward {
		return true
	}
	return false
}
