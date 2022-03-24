package ports

//go:generate mockgen -destination=../mocks/mocks.go -package=mocks -source=./ports.go

import "gosanta/internal"

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
	AssignPhishingAward(cpa *CreatePhishingAward)
}

type AwardReadingService interface {
	Get(id string) (awards.PhishingAward, error)
	GetUserAwards(uId awards.UserId) ([]awards.PhishingAward, error)
	GetCompanyAwards(cId awards.CompanyId) ([]awards.PhishingAward, error)
}
