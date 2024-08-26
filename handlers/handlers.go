package hanlders

import (
	"distro-hub/domain"
)

type Handler struct {
	auth *domain.AuthService
}

func NewHandler() *Handler {
	return &Handler{}
}
