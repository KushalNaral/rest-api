package auth

import (
	"database/sql"
	"fmt"
	"rest-api/internal/config"
	"rest-api/internal/queue"

	"github.com/thecodearcher/limen"
	sqladapter "github.com/thecodearcher/limen/adapters/sql"
	credentialpassword "github.com/thecodearcher/limen/plugins/credential-password"
)

type AuthPool struct {
	Auth *limen.Limen
}

func NewAuth(cfg *config.Config, db *sql.DB, rmq *queue.RabbitMQ) (*AuthPool, error) {

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
		Email: limen.NewDefaultEmailConfig(
			limen.WithEmailVerification(
				limen.WithSendEmailVerificationMail(func(email, token string) {
					// send email with verification link

					emailVerificationPayload := queue.EmailVerificationPayload{
						Email: email,
						Token: token,
					}

					if rmq != nil {
						err := rmq.PublishEmailVerification(emailVerificationPayload)
						if err != nil {
							fmt.Println("error publishing email verification event:", err)
						}
					} else {
						// Fallback if no RabbitMQ configured
						fmt.Printf("[auth] verification token for %s: %s\n", email, token)
					}
				}),
			),
		),
	}

	auth, err := limen.New(config)
	if err != nil {
		return nil, err
	}

	return &AuthPool{
		Auth: auth,
	}, nil
}
