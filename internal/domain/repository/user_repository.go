package repository

import "distro-hub/internal/domain/entity"

type UserRepository interface {
	GetByID(id string) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id string) error
	List() ([]*entity.User, error)
}
