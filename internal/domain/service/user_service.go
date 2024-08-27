package service

import "distro-hub/internal/domain/entity"

type UserService interface {
	GetUser(id string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	DeleteUser(id string) error
	ListUsers() ([]*entity.User, error)
	AuthenticateUser(username, password string) (*entity.User, error)
	ChangePassword(userID, oldPassword, newPassword string) error
	ValidateUser(user *entity.User) error
}
