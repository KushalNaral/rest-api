package api

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/config"
	"time"
)

var BaseUrl = "/api/"

func RegisterBaseRoutes(cfg *config.Config) http.Handler {

	root := http.NewServeMux()
	apiVersion := cfg.ApiVersion

	v1 := http.NewServeMux()
	registerHealthRoutes(v1)

	root.Handle(BaseUrl+apiVersion+"/", http.StripPrefix(BaseUrl+apiVersion, v1))

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
