package main

import (
	"fmt"
	awards "gosanta/internal"
	"gosanta/internal/awarding"
	"gosanta/internal/awssqs"
	"gosanta/internal/eventbroker"
	"gosanta/internal/eventlogging"
	"gosanta/internal/eventpublishing"
	"gosanta/internal/postgres"
	"gosanta/internal/ranking"
	"gosanta/internal/rest"
	"gosanta/internal/usernotifying"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	fmt.Print("Starting awarding job.")
	err := run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func run() error {
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

	// services
	awardSrvc := awarding.NewAwardService(ar, ur, per)
	eventLogger := eventlogging.New(per, &client)
	rankingSrvc := ranking.NewService(ar, ur)
	restSrv := rest.New(&rankingSrvc)

	// stubs for now
	usrNotifyer := usernotifying.New()
	eventPublisher := eventpublishing.New()

	awarder := eventbroker.NewAwarderNotifier(&eventLogger, &awardSrvc, &usrNotifyer, &eventPublisher)

	eventChan := make(chan awards.UserPhishingEvent)
	awardChan := make(chan awards.UserAwardEvent)
	awarder.Start(eventChan, awardChan)

	err := http.ListenAndServe(":"+config.http_port, &restSrv)

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
