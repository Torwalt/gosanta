package main

import (
	"fmt"
	"gosanta/internal/awssqs"
	"gosanta/internal/eventlogging"
	"gosanta/internal/postgres"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	fmt.Print("Starting event sync.")
	err := run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Print("Events loaded without issue.")
}

func run() error {
	sess := session.New()
	s := sqs.New(sess)
	config := loadFromEnv()
	er := awssqs.New(s, config.queueURL)

	pconf := postgres.Config{
		Host:   config.postgres_host,
		Port:   config.postgres_port,
		User:   config.postgres_user,
		Secret: config.postgres_secret,
		Name:   config.postgres_name,
	}
	sqldb := postgres.NewDb(pconf)
	erRepo := postgres.NewPhishingEventRepository(sqldb)
	el := eventlogging.New(erRepo, &er)

	err := el.LogNewEvents()
	if err != nil {
		return fmt.Errorf("error when logging new events: %v", err)
	}

	return nil
}

type config struct {
	queueURL string

	postgres_host   string
	postgres_port   string
	postgres_user   string
	postgres_secret string
	postgres_name   string
}

func loadFromEnv() *config {
	return &config{
		queueURL:        os.Getenv("queueURL"),
		postgres_host:   os.Getenv("POSTGRES_HOST"),
		postgres_port:   os.Getenv("POSTGRES_PORT"),
		postgres_user:   os.Getenv("POSTGRES_USER"),
		postgres_secret: os.Getenv("POSTGRES_SECRET"),
		postgres_name:   os.Getenv("POSTGRES_NAME"),
	}
}
