package repository

import "distro-hub/internal/domain/entity"

type DistroRepository interface {
	GetByID(id string) (*entity.Distro, error)
	Create(distro *entity.Distro) error
	Update(distro *entity.Distro) error
	Delete(id string) error
	List() ([]*entity.Distro, error)
	FindByName(name string) (*entity.Distro, error)
}
