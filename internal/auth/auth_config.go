package auth

import (
	"database/sql"

	"github.com/thecodearcher/limen"
	sqladapter "github.com/thecodearcher/limen/adapters/sql"
)

type AuthPool struct {
	Auth *limen.Limen
}

func NewAuth(db *sql.DB) (*AuthPool, error) {

	config := &limen.Config{
		BaseURL:  "http://localhost:8080",
		Database: sqladapter.NewPostgreSQL(db),
		CLI: &limen.CLIConfig{
			Enabled: true,
		},
	}

	auth, err := limen.New(config)
	if err != nil {
		return nil, err
	}

	return &AuthPool{
		Auth: auth,
	}, nil
}
