package repository

import (
	"database/sql"
	"employeeleave/model"
)

type RoleRepository interface {
	Create(payload model.Role) error
	GetRole(roleName string) (model.Role, error)
	List() ([]model.Role, error)
}

type roleRepository struct {
	db *sql.DB
}

func (r *roleRepository) Create(payload model.Role) error {
	_, err := r.db.Exec("INSERT INTO role (id, role_name) VALUES ($1, $2)", payload.Id, payload.RoleName)
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) GetRole(roleName string) (model.Role, error) {
	var user model.Role
	err := r.db.QueryRow("SELECT id, role_name FROM role WHERE role_name = $1", roleName).Scan(&user.Id, &user.RoleName)
	if err != nil {
		return model.Role{}, err
	}
	return user, nil
}

func (r *roleRepository) List() ([]model.Role, error) {
	var roles []model.Role

	rows, err := r.db.Query("SELECT id, role_name FROM role")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var role model.Role
		err := rows.Scan(&role.Id, &role.RoleName)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{db: db}
}