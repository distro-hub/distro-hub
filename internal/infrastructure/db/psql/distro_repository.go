package psql

import (
	"database/sql"
	"distro-hub/internal/domain/entity"
)

type distroRepository struct {
	db *sql.DB
}

func NewDistroRepository(db *sql.DB) *distroRepository {
	return &distroRepository{
		db: db,
	}
}

func (r *distroRepository) GetByID(id string) (*entity.Distro, error) {
	var d entity.Distro
	err := r.db.QueryRow("SELECT id, name, description FROM distros WHERE id = $1", id).
		Scan(&d.ID, &d.Name, &d.Description)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *distroRepository) Create(distro *entity.Distro) error {
	_, err := r.db.Exec("INSERT INTO distros (id, name, description) VALUES ($1, $2, $3)",
		distro.ID, distro.Name, distro.Description)
	return err
}

func (r *distroRepository) Update(distro *entity.Distro) error {
	return nil
}

func (r *distroRepository) Delete(id string) error {
	return nil
}

func (r *distroRepository) List() ([]*entity.Distro, error) {
	distros := []*entity.Distro{}
	err := r.db.QueryRow("SELECT * FROM distros").Scan(&distros)
	if err != nil {
		return nil, err
	}
	return distros, nil
}

func (r *distroRepository) FindByName(name string) (*entity.Distro, error) {
	distro := *&entity.Distro{}
	return &distro, nil
}
