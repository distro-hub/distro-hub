package hanlders

import (
	"context"
	"distro-hub/data"
	"distro-hub/domain"
	"distro-hub/views"
	"net/http"
)

func (h *Handler) DistroMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HomePage)
	mux.HandleFunc("POST /distro", h.NewDistro)

	return http.StripPrefix("/distro", mux)
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	title := "Distro Hub"
	ctx := context.WithValue(r.Context(), "currentUser", data.User)
	views.HomePage(title, data.Distros).Render(ctx, w)
}

var id = 3

func (h *Handler) NewDistro(w http.ResponseWriter, r *http.Request) {
	id = id + 1
	newDistro := domain.Distro{ID: id, Name: r.FormValue("distro_name")}
	data.Distros = append(data.Distros, newDistro)
	views.DistroList(data.Distros).Render(r.Context(), w)
}
