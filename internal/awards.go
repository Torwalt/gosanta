package awards

import (
	"fmt"
	"time"
)

type UserId int64
type CompanyId int64

// The reason an award was given.
type AwardType int

const (
	OpenAward AwardType = iota
	ReportAward
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
	Reason     AwardType
	EmailRef   string
}

// A user that can receive an award.
type User struct {
	Id        UserId
	CompanyId CompanyId
	Awards    []PhishingAward
}

// A company to which users belong.
type Company struct {
	id    CompanyId
	users []User
}

// Prepare a new award for a user. An award cannot be gained twice.
func New(u User, emailRef string, reason AwardType) (*PhishingAward, error) {
	if isDuplicate(u, emailRef, reason) {
		return nil, &Error{Code: DuplicateError, Err: fmt.Errorf(
			"user %v has already earned award from email %v", u.Id, emailRef)}
	}
	return &PhishingAward{
		AssignedTo: u.Id,
		EarnedOn:   time.Now(),
		Reason:     reason,
		EmailRef:   emailRef,
	}, nil
}

func isDuplicate(u User, emailRef string, reason AwardType) bool {
	for _, a := range u.Awards {
		if a.EmailRef == emailRef && a.Reason == reason {
			return true
		}
	}
	return false
}
