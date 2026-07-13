package auth

import (
	"database/sql"
	"rest-api/internal/config"

	"github.com/thecodearcher/limen"
	sqladapter "github.com/thecodearcher/limen/adapters/sql"
	credentialpassword "github.com/thecodearcher/limen/plugins/credential-password"
)

type AuthPool struct {
	Auth *limen.Limen
}

func NewAuth(cfg *config.Config, db *sql.DB) (*AuthPool, error) {

	config := &limen.Config{
		BaseURL:  "http://localhost:8080",
		Database: sqladapter.NewPostgreSQL(db),
		CLI: &limen.CLIConfig{
			Enabled: true,
		},
		HTTP: limen.NewDefaultHTTPConfig(
			limen.WithHTTPBasePath("/api/" + cfg.ApiVersion + "/auth"),
		),
		Plugins: []limen.Plugin{credentialpassword.New(
			credentialpassword.WithAutoSignInOnSignUp(false),
			credentialpassword.WithPasswordMinLength(8),
			credentialpassword.WithPasswordRequireSymbols(true),
		)},
	}

	auth, err := limen.New(config)
	if err != nil {
		return nil, err
	}

	return &AuthPool{
		Auth: auth,
	}, nil
}
