package service

import "distro-hub/internal/domain/entity"

type DistroService interface {
	GetDistro(id string) (*entity.Distro, error)
	CreateDistro(distro *entity.Distro) error
	UpdateDistro(distro *entity.Distro) error
	DeleteDistro(id string) error
	ListDistros() ([]*entity.Distro, error)
	SearchDistroByName(name string) (*entity.Distro, error)
	ValidateDistro(distro *entity.Distro) error
}
