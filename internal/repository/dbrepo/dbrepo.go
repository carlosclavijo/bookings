package dbrepo

import (
	"database/sql"

	"github.com/carlosclavijo/bookings/internal/config"
	"github.com/carlosclavijo/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
