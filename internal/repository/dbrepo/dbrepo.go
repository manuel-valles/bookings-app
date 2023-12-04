package dbrepo

import (
	"database/sql"

	"github.com/manuel-valles/bookings-app.git/internal/config"
	"github.com/manuel-valles/bookings-app.git/internal/repository"
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
