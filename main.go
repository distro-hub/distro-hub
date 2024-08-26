package main

import (
	hanlders "distro-hub/handlers"
	"distro-hub/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	h := hanlders.NewHandler()
	// auth
	mux.Handle("/auth/", h.AuthMux())

	mux.HandleFunc("/", h.HomePage)

	// distro stuff
	mux.Handle("/distro/", h.DistroMux())

	s := &http.Server{
		Addr:    ":6969",
		Handler: middleware.Auth(mux),
	}

	s.ListenAndServe()
}
