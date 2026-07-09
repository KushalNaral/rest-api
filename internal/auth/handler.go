package auth

import "net/http"

func RegisterAuthHandler(handler http.Handler) http.Handler {

	mux := http.NewServeMux()
	mux.Handle("/auth/", handler)
	return mux
}
