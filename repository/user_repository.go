package repository

import (
	"database/sql"
	"employeeleave/model"
)

type UserRepository interface {
	Create(payload model.UserCredential) error
}

type userRepository struct {
	db *sql.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(payload model.UserCredential) error {
	_, err := u.db.Exec("INSERT INTO user_credential (id, username, password, role_id) VALUES ($1, $2, $3, $4)", payload.ID, payload.Username, payload.Password, payload.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
