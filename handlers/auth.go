package hanlders

import (
	"distro-hub/views"
	"distro-hub/views/auth"
	"fmt"
	"net/http"
)

func (h *Handler) AuthMux() http.Handler {
	mux := http.NewServeMux()

	// handle pages
	mux.HandleFunc("/", h.HandleAuth)
	mux.HandleFunc("/login", h.HandleLoginTab)
	mux.HandleFunc("/register", h.HandleRegisterTab)

	// handle data
	mux.HandleFunc("POST /login", h.HandleLogin)

	return http.StripPrefix("/auth", mux)
}

func (h *Handler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	views.Auth().Render(r.Context(), w)
}
func (h *Handler) HandleLoginTab(w http.ResponseWriter, r *http.Request) {
	auth.LoginTab().Render(r.Context(), w)
}
func (h *Handler) HandleRegisterTab(w http.ResponseWriter, r *http.Request) {
	auth.RegisterTab().Render(r.Context(), w)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(email, password)
	http.Redirect(w, r, "/", 301)
}
