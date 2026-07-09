package db

import (
	"database/sql"
	"fmt"
	"rest-api/internal/config"
)

type Database struct {
	Db *sql.DB
}

func NewConn(cfg *config.Config) (*Database, error) {

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &Database{
		Db: db,
	}, nil
}
