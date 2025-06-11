package awarding

import (
	"time"

	awards "gosanta/internal"
	"gosanta/internal/ports"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	next   ports.AwardAssigner
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s ports.AwardAssigner) ports.AwardAssigner {
	return &loggingService{logger, s}
}

func (s *loggingService) AssignAward(
	event awards.UserPhishingEvent,
) (usrAwardEvent awards.UserAwardEvent, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_award",
			"event", event,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.next.AssignAward(event)
}
