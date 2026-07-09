package db

import (
	"database/sql"
	"fmt"
	"rest-api/internal/config"

	"github.com/pressly/goose/v3"
)

var gooseUp = goose.Up

func Migrate(cfg *config.Config, db *sql.DB) error {
	tableName := cfg.GooseTable
	if tableName == "" {
		tableName = "migrations"
	}
	goose.SetTableName(tableName)
	if err := gooseUp(db, cfg.GooseMigrationDir); err != nil {
		return fmt.Errorf("migrating database: %w", err)
	}

	return nil
}
