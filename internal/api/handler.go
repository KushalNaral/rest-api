package api

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/auth"
	"rest-api/internal/config"
	"time"
)

var BaseUrl = "/api/"

func RegisterBaseRoutes(cfg *config.Config, pool *auth.AuthPool) http.Handler {
	root := http.NewServeMux()

	apiPrefix := BaseUrl + cfg.ApiVersion

	api := registerAPIRoutes(cfg)
	authAPI := auth.RegisterAuthRoutes(cfg, pool)

	root.HandleFunc("/", handleBaseResponse(cfg))
	root.HandleFunc(apiPrefix, handleBaseResponse(cfg))

	root.Handle(apiPrefix+"/auth/", authAPI)
	root.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, api))

	return root
}

type HealthResponse struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func registerAPIRoutes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	registerHealthRoutes(mux)
	return mux
}

func registerHealthRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := HealthResponse{
			Status:    http.StatusOK,
			Timestamp: time.Now(),
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func handleBaseResponse(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := struct {
			APIVersion string `json:"apiVersion"`
			Message    string `json:"message"`
		}{
			APIVersion: cfg.ApiVersion,
			Message:    "OK : Connection Successful",
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}
