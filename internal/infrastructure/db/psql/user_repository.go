package psql

import (
	"database/sql"
	"distro-hub/internal/domain/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, email, password FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, email, password FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *entity.User) error {
	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user *entity.User) error {
	query := "UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4"
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List() ([]*entity.User, error) {
	query := "SELECT id, username, email, password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
