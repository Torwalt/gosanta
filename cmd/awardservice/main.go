package main

import (
	"gosanta/internal/postgres"
	"gosanta/internal/ranking"
	"gosanta/internal/server"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
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

	awardRepo := postgres.NewAwardRepository(sqldb)
	userRepo := postgres.NewUserRepository(sqldb)

	r := ranking.New(awardRepo, userRepo)

	srv := server.New(&r)

	http.ListenAndServe(":"+config.http_port, &srv)

	return nil
}

type config struct {
	http_port string

	postgres_host   string
	postgres_port   string
	postgres_user   string
	postgres_secret string
	postgres_name   string
}

func loadFromEnv() *config {
	return &config{
		http_port:       os.Getenv("HTTP_PORT"),
		postgres_host:   os.Getenv("POSTGRES_HOST"),
		postgres_port:   os.Getenv("POSTGRES_PORT"),
		postgres_user:   os.Getenv("POSTGRES_USER"),
		postgres_secret: os.Getenv("POSTGRES_SECRET"),
		postgres_name:   os.Getenv("POSTGRES_NAME"),
	}
}
