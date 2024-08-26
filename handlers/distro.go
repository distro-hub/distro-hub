package hanlders

import (
	"context"
	"distro-hub/domain"
	"distro-hub/views"
	"net/http"
)

var distros = domain.Distros{
	{ID: 1, Name: "Ubuntu"},
	{ID: 2, Name: "Debian"},
	{ID: 3, Name: "Arch"},
}

var user = domain.User{
	ID:       1,
	Name:     "Radoje",
	LastName: "Jezdic",
	Age:      27,
}

func (h *Handler) DistroMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HomePage)

	return http.StripPrefix("/distro", mux)
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	title := "Distro Hub"
	ctx := context.WithValue(r.Context(), "currentUser", user)
	views.HomePage(title, distros).Render(ctx, w)
}
