package data

import "distro-hub/domain"

var Distros = domain.Distros{
	{ID: 1, Name: "Ubuntu"},
	{ID: 2, Name: "Debian"},
	{ID: 3, Name: "Arch"},
}

var User = domain.User{
	ID:       1,
	Name:     "Radoje",
	LastName: "Jezdic",
	Age:      27,
}
