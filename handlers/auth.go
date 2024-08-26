package hanlders

import (
	"distro-hub/views"
	"distro-hub/views/auth"
	"net/http"
)

func (h *Handler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	views.Auth().Render(r.Context(), w)
}

func (h *Handler) HandleLoginTab(w http.ResponseWriter, r *http.Request) {
	auth.LoginTab().Render(r.Context(), w)
}
func (h *Handler) HandleRegisterTab(w http.ResponseWriter, r *http.Request) {
	auth.RegisterTab().Render(r.Context(), w)
}

func (h *Handler) AuthMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HandleAuth)
	mux.HandleFunc("/login", h.HandleLoginTab)
	mux.HandleFunc("/register", h.HandleRegisterTab)

	return http.StripPrefix("/auth", mux)
}
