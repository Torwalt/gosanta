package postgres

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host   string
	Port   string
	User   string
	Secret string
	Name   string
}

func NewDb(config Config) *sql.DB {
	pgDSN := fmt.Sprintf(
		DSN,
		config.User,
		config.Secret,
		config.Host,
		config.Port,
		config.Name,
	)
	return sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pgDSN)))
}
