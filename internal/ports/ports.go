package ports

//go:generate mockgen -destination=../mocks/ports.go -package=mocks -source=./ports.go

import (
	"gosanta/internal"
	"gosanta/pkg"
)

type AwardReadRepository interface {
	Get(id int64) (*awards.PhishingAward, error)
	GetUserAwards(id awards.UserId) ([]awards.PhishingAward, error)
	GetByEmailRef(id awards.UserId, ref string) (*awards.PhishingAward, error)
}

type AwardRepository interface {
	AwardReadRepository
	Add(award *awards.PhishingAward) error
	UpdateExisting(existing, award *awards.PhishingAward) error
	Delete(id int64) error
}

type UserReadRepository interface {
	Get(uId awards.UserId) (*awards.User, error)
	GetCompanyUsers(cId awards.CompanyId) ([]awards.User, error)
}

type AwardReadingService interface {
	Get(id int64) (*awards.PhishingAward, error)
	GetUserAwards(uId awards.UserId) ([]awards.PhishingAward, error)
	CalcLeaderboard(uId awards.UserId) (*awards.Leaderboard, error)
}

type EventReadRepository interface {
	GetUnprocessed() ([]awards.UserPhishingEvent, error)
	ClickedExists(uID awards.UserId, emailRef string) (bool, error)
}

type EventRepositoryProcessor interface {
	EventReadRepository
	MarkAsProcessed(event *awards.UserPhishingEvent) error
}

type EventRepository interface {
	EventReadRepository
	Write(upe awards.UserPhishingEvent) error
}

type EventQueue interface {
	GetNextMessages() ([]events.PhishingEvent, error)
	DeleteMessage(eventID string) error
}

type MailSender interface {
	SendToUser(awards.UserAwardEvent) error
}

type EventPublisher interface {
	PublishEvent(awards.UserAwardEvent) error
}

type EventLogReader interface {
	LogNewEvents() ([]awards.UserPhishingEvent, error)
	GetUnprocessedEvents() ([]awards.UserPhishingEvent, error)
}

type AwardAssigner interface {
	AssignAward(event awards.UserPhishingEvent) (awards.UserAwardEvent, error)
}
