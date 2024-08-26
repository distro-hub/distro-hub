package main

import (
	"context"
	"distro-hub/domain"
	hanlders "distro-hub/handlers"
	"distro-hub/middleware"
	"distro-hub/views"
	"fmt"
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
		{ID: 1, Name: "Ubuntu"},
		{ID: 2, Name: "Debian"},
		{ID: 3, Name: "Arch"},
	}
	id := 3

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "currentUser", user)
		views.HomePage("Distro Hub", distros).Render(ctx, w)
	})

	mux.HandleFunc("POST /distro", func(w http.ResponseWriter, r *http.Request) {
		id = id + 1
		newDistro := domain.Distro{
			ID:   id,
			Name: r.FormValue("distro_name"),
		}
		fmt.Println(id)
		distros = append(distros, newDistro)
		views.DistroList(distros).Render(r.Context(), w)
	})
	mux.HandleFunc("POST /auth/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		pass := r.FormValue("password")

		fmt.Println(email, pass)

		http.Redirect(w, r, "/", 301)
		return
	})

	h := hanlders.NewHandler()
	// auth
	mux.Handle("/auth/", h.AuthMux())

	s := &http.Server{
		Addr:    ":6969",
		Handler: middleware.Auth(mux),
	}

	s.ListenAndServe()
}
