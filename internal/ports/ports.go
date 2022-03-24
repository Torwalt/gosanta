package ports

//go:generate mockgen -destination=../mocks/ports.go -package=mocks -source=./ports.go

import (
	"gosanta/internal"
	"gosanta/pkg"
)

type CreatePhishingAward struct {
	Id       awards.UserId
	EmailRef string
	Action   awards.PhishingAction
}

type AwardReadRepository interface {
	Get(id int64) (*awards.PhishingAward, error)
	GetByUserId(id awards.UserId) ([]awards.PhishingAward, error)
}

type AwardRepository interface {
	AwardReadRepository
	Add(a *awards.PhishingAward) (*awards.PhishingAward, error)
	Delete(id int) error
}

type UserReadRepository interface {
	Get(uId awards.UserId) (*awards.User, error)
}

type AwardAssigningService interface {
	HandlePhishingEvent(cpa *CreatePhishingAward)
}

type AwardReadingService interface {
	Get(id string) (awards.PhishingAward, error)
	GetUserAwards(uId awards.UserId) ([]awards.PhishingAward, error)
	GetCompanyAwards(cId awards.CompanyId) ([]awards.PhishingAward, error)
}

type EventReadRepository interface {
	GetUnprocessed() ([]awards.UserPhishingEvent, error)
}

type EventRepository interface {
	EventReadRepository
	Write(upe awards.UserPhishingEvent) error
}

type EventQueue interface {
	GetNextMessages() ([]events.PhishingEvent, error)
	DeleteMessage(eventID string) error
}
