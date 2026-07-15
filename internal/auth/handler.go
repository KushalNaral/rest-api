package auth

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/config"
)

func RegisterAuthRoutes(cfg *config.Config, pool *AuthPool) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", pool.Auth.Handler())

	mux.HandleFunc("/me", handleProfile(pool))

	return mux
}

func handleProfile(pool *AuthPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		session, err := pool.Auth.GetSession(r)
		if err != nil {
			http.Error(w, "error : not authenticated", http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]any{"user": session.User.Raw()})
	}
}
