package main

import (
	"database/sql"
	"fmt"
	"gosanta/internal/postgres"
	"gosanta/internal/ranking"
	"gosanta/internal/server"
	"net/http"

	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	run()
}

func run() error {
	pgDSN := fmt.Sprintf(postgres.DSN)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pgDSN)))

	awardRepo := postgres.NewAwardRepository(sqldb)
	userRepo := postgres.NewUserRepository(sqldb)

	r := ranking.New(awardRepo, userRepo)

	srv := server.New(&r)

	http.ListenAndServe(":8080", &srv)

	return nil
}
