package main

import (
	"context"
	"distro-hub/domain"
	"distro-hub/views"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	user := domain.User{
		ID:       1,
		Name:     "Radoje",
		LastName: "Jezdic",
		Age:      27,
	}

	distros := domain.Distros{
		{
			ID:   1,
			Name: "Ubuntu",
		},
		{
			ID:   2,
			Name: "Debian",
		},
		{
			ID:   3,
			Name: "Arch",
		},
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "currentUser", user)
		views.HomePage("Distro Hub", distros).Render(ctx, w)
	})

	s := &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}

	s.ListenAndServe()
}
