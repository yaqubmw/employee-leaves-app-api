package repository

import (
	"database/sql"
	"employeeleave/model"
)

type RoleRepository interface {
	Create(payload model.Role) error
	Get(id string) (model.Role, error)
	GetByName(roleName string) (model.Role, error)
	List() ([]model.Role, error)
	Update(payload model.Role) error
	Delete(id string) error
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

func (r *roleRepository) Get(id string) (model.Role, error) {
	var role model.Role
	err := r.db.QueryRow("SELECT id, role_name FROM role WHERE id=$1", id).Scan(&role.Id, &role.RoleName)
	if err != nil {
		return model.Role{}, err
	}
	return role, nil
}

func (r *roleRepository) GetByName(roleName string) (model.Role, error) {
	var role model.Role
	err := r.db.QueryRow("SELECT id, role_name FROM role WHERE role_name ILIKE $1", "%"+roleName+"%").Scan(&role.Id, &role.RoleName)
	if err != nil {
		return model.Role{}, err
	}
	return role, nil
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

func (r *roleRepository) Update(payload model.Role) error {
	_, err := r.db.Exec("UPDATE role SET role_name=$1 WHERE id=$2", payload.RoleName, payload.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r *roleRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM role WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{db: db}
}