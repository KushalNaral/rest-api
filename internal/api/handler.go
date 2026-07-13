package api

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/config"
	"time"
)

var BaseUrl = "/api/"

func RegisterBaseRoutes(cfg *config.Config, authHandler http.Handler) http.Handler {
	root := http.NewServeMux()
	api := http.NewServeMux()

	registerHealthRoutes(api)

	apiPrefix := BaseUrl + cfg.ApiVersion

	root.HandleFunc("/", handleBaseResponse(cfg))
	root.HandleFunc(apiPrefix, handleBaseResponse(cfg))
	root.Handle(apiPrefix+"/auth/", authHandler)

	root.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, api))

	return root
}

type HealthResponse struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
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
