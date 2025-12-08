package database

import (
	"cinema_service/internal/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	dsn := cfg.Postgres.ConnectionURL()

	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("sqlx open error: %w", err)
	}

	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConnections)
	db.SetConnMaxLifetime(cfg.Postgres.ConnectionMaxLifetime * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	return db, nil
}
