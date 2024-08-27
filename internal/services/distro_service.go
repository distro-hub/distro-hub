package services

import (
	"distro-hub/internal/domain/entity"
	"distro-hub/internal/domain/repository"
	"distro-hub/internal/domain/service"
)

type distroService struct {
	repo repository.DistroRepository
}

func NewDistroService(repo repository.DistroRepository) service.DistroService {
	return &distroService{repo: repo}
}

func (s *distroService) GetDistro(id string) (*entity.Distro, error) {
	return s.repo.GetByID(id)
}

func (s *distroService) CreateDistro(distro *entity.Distro) error {
	if err := s.ValidateDistro(distro); err != nil {
		return err
	}
	return s.repo.Create(distro)
}

func (s *distroService) UpdateDistro(distro *entity.Distro) error {
	return nil
}

func (s *distroService) DeleteDistro(id string) error {
	return nil
}

func (s *distroService) ListDistros() ([]*entity.Distro, error) {
	return s.repo.List()
}

func (s *distroService) SearchDistroByName(name string) (*entity.Distro, error) {
	return s.repo.FindByName(name)
}

func (s *distroService) ValidateDistro(distro *entity.Distro) error {
	// Implement validation logic here
	return nil
}
