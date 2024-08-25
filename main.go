package main

import (
	"context"
	"distro-hub/domain"
	"distro-hub/middleware"
	"distro-hub/views"
	"distro-hub/views/auth"
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

	mux.HandleFunc("GET /auth", func(w http.ResponseWriter, r *http.Request) {
		views.Auth().Render(r.Context(), w)
	})
	mux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginTab().Render(r.Context(), w)
	})
	mux.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		auth.RegisterTab().Render(r.Context(), w)
	})
	mux.HandleFunc("POST /auth/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		pass := r.FormValue("password")

		fmt.Println(email, pass)

		http.Redirect(w, r, "/", 301)
		return
	})

	s := &http.Server{
		Addr:    ":6969",
		Handler: middleware.Auth(mux),
	}

	s.ListenAndServe()
}
