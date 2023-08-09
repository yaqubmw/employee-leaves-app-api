package repository

import (
	"database/sql"
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/utils/common"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(payload model.UserCredential) error
	GetUsername(username string) (model.UserCredential, error)
	GetUsernamePassword(username, password string) (model.UserCredential, error)
	BaseRepositoryPaging[model.UserCredential]
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

// GetUsername implements UserRepository.
func (u *userRepository) GetUsername(username string) (model.UserCredential, error) {
	var user model.UserCredential
	err := u.db.QueryRow("SELECT id, username, password from user_credential WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return model.UserCredential{}, err
	}
	return user, nil
}

// GetUsernamePassword implements UserRepository.
func (u *userRepository) GetUsernamePassword(username string, password string) (model.UserCredential, error) {
	user, err := u.GetUsername(username)
	if err != nil {
		return model.UserCredential{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("failed to verify password hash: %v", err)
	}
	return user, nil
}

//Paging implement user repository
func (u *userRepository) Paging(requestPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)
	rows, err := u.db.Query("SELECT id, username, role_id FROM user_credential LIMIT $1 OFFSET $2", paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	var users []model.UserCredential
	for rows.Next() {
		var user model.UserCredential
		err := rows.Scan(&user.ID, &user.Username, &user.RoleId)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		users = append(users, user)
	}

	//count users
	var totalRows int
	row := u.db.QueryRow("SELECT COUNT(*) FROM user")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return users, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
