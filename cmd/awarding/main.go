package main

import (
	"fmt"
	"gosanta/internal/awarding"
	"gosanta/internal/postgres"
	"log"
	"os"
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
	er := postgres.NewPhishingEventRepository(sqldb)

	awardSrv := awarding.NewAwardService(ar, ur, er)

	err := awardSrv.ProcessPhishingEvents()
	if err != nil {
		return fmt.Errorf("an error occurred when assigning phishing awards: %v", err)
	}
	return nil
}

type config struct {
	postgres_host   string
	postgres_port   string
	postgres_user   string
	postgres_secret string
	postgres_name   string
}

func loadFromEnv() *config {
	return &config{
		postgres_host:   os.Getenv("POSTGRES_HOST"),
		postgres_port:   os.Getenv("POSTGRES_PORT"),
		postgres_user:   os.Getenv("POSTGRES_USER"),
		postgres_secret: os.Getenv("POSTGRES_SECRET"),
		postgres_name:   os.Getenv("POSTGRES_NAME"),
	}
}
