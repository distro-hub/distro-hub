package handler

import (
	"context"
	"distro-hub/data"
	"distro-hub/internal/domain/service"
	"distro-hub/internal/interface/web/templates"
	"distro-hub/internal/interface/web/templates/auth"
	"net/http"
)

type WebHandler struct {
	distroService service.DistroService
	userService   service.UserService
}

func NewWebHandler(ds service.DistroService, us service.UserService) *WebHandler {
	return &WebHandler{
		distroService: ds,
		userService:   us,
	}
}

func (h *WebHandler) Home(w http.ResponseWriter, r *http.Request) {
	// distros, err := h.distroService.ListDistros()
	// if err != nil {
	// 	http.Error(w, "Failed to fetch distros", http.StatusInternalServerError)
	// 	return
	// }
	ctx := context.WithValue(r.Context(), "currentUser", data.User)
	distros := data.Distros // dammy data
	templates.Home(distros).Render(ctx, w)
}

func (h *WebHandler) Auth(w http.ResponseWriter, r *http.Request) {
	templates.Auth().Render(r.Context(), w)
}
func (h *WebHandler) LoginForm(w http.ResponseWriter, r *http.Request) {
	auth.LoginForm().Render(r.Context(), w)
}
func (h *WebHandler) RegisterForm(w http.ResponseWriter, r *http.Request) {
	auth.RegisterForm().Render(r.Context(), w)
}
