package main

import (
	"fmt"
	"net/http"
	"os"

	awards "gosanta/internal"
	"gosanta/internal/awarding"
	"gosanta/internal/awssqs"
	"gosanta/internal/eventbroker"
	"gosanta/internal/eventlogging"
	"gosanta/internal/eventpublishing"
	"gosanta/internal/ports"
	"gosanta/internal/postgres"
	"gosanta/internal/ranking"
	"gosanta/internal/rest"
	"gosanta/internal/usernotifying"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/go-kit/kit/log"
)

func main() {
	fmt.Print("Starting awarding job.")

	// logging
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)

	err := run(logger)
	if err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}
}

func run(logger log.Logger) error {
	config := loadFromEnv()

	// repositories
	pconf := postgres.Config{
		Host:   config.postgres_host,
		Port:   config.postgres_port,
		User:   config.postgres_user,
		Secret: config.postgres_secret,
		Name:   config.postgres_name,
	}
	sqldb := postgres.NewDb(pconf)
	ar := postgres.NewAwardRepository(sqldb)
	ur := postgres.NewUserRepository(sqldb)
	per := postgres.NewPhishingEventRepository(sqldb)

	// sqs client
	sess := session.New()
	s := sqs.New(sess)
	client := awssqs.New(s, config.queueURL)

	// awarding service
	var awardSrvc ports.AwardAssigner
	awardSrvc = awarding.NewAwardService(ar, ur, per)
	awardSrvc = awarding.NewLoggingService(log.With(logger, "component", "awarding"), awardSrvc)

	// event retrieval and persistence service
	eventLogger := eventlogging.New(per, &client)

	// award ranking and retrieval service
	rankingSrvc := ranking.NewService(ar, ur)

	// rest server
	restSrv := rest.New(&rankingSrvc)

	// stubs for now
	usrNotifyer := usernotifying.New()
	eventPublisher := eventpublishing.New()

	// awarding flow orchestrator
	eventChan := make(chan awards.UserPhishingEvent)
	awardChan := make(chan awards.UserAwardEvent)
	awarder := eventbroker.NewAwarderNotifier(
		&eventLogger,
		awardSrvc,
		&usrNotifyer,
		&eventPublisher,
		log.With(logger, "component", "award-notifier"),
		eventChan,
		awardChan,
	)

	err := awarder.Start()
	if err != nil {
		return fmt.Errorf("could not start awarder: %v", err)
	}

	err = http.ListenAndServe(":"+config.http_port, &restSrv)
	if err != nil {
		return fmt.Errorf("could not start server: %v", err)
	}
	return nil
}

type config struct {
	http_port string

	queueURL string

	postgres_host   string
	postgres_port   string
	postgres_user   string
	postgres_secret string
	postgres_name   string
}

func loadFromEnv() *config {
	return &config{
		http_port:       os.Getenv("HTTP_PORT"),
		queueURL:        os.Getenv("queueURL"),
		postgres_host:   os.Getenv("POSTGRES_HOST"),
		postgres_port:   os.Getenv("POSTGRES_PORT"),
		postgres_user:   os.Getenv("POSTGRES_USER"),
		postgres_secret: os.Getenv("POSTGRES_SECRET"),
		postgres_name:   os.Getenv("POSTGRES_NAME"),
	}
}
