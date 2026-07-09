package db

import (
	"database/sql"
	"rest-api/internal/config"
	"testing"

	"github.com/pressly/goose/v3"
)

func TestMigrate(t *testing.T) {
	old := gooseUp
	defer func() { gooseUp = old }()

	called := false

	gooseUp = func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error {
		called = true

		if dir != "migrations" {
			t.Fatalf("expected migrations, got %q", dir)
		}

		return nil
	}

	cfg := &config.Config{
		GooseMigrationDir: "migrations",
		GooseTable:        "test_table",
	}

	if err := Migrate(cfg, &sql.DB{}); err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("expected goose.Up to be called")
	}
}
