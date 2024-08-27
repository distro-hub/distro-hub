package data

import (
	"distro-hub/internal/domain/entity"
)

var Distros = []entity.Distro{
	{ID: 1, Name: "Ubuntu"},
	{ID: 2, Name: "Debian"},
	{ID: 3, Name: "Arch"},
}

var User = entity.User{
	ID:       "1",
	Username: "Radoje",
	Email:    "john@doe.com",
	Password: "123",
}
