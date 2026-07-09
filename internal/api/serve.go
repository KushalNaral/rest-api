package api

import (
	"fmt"
	"net/http"
	"rest-api/internal/config"
)

func Serve(cfg *config.Config, handler http.Handler) error {
	addr := ":" + cfg.ApiPort

	fmt.Printf("Server listening on http://localhost%s\n", addr)

	return http.ListenAndServe(addr, handler)
}
