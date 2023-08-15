package repository

import (
	"employeeleave/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(payload model.Role) error
	Get(id string) (model.Role, error)
	GetRoleByName(roleName string) (model.Role, error)
	List() ([]model.Role, error)
	Update(payload model.Role) error
	Delete(id string) error
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) Create(payload model.Role) error {
	return r.db.Create(&payload).Error
}

func (r *roleRepository) Get(id string) (model.Role, error) {
	var role model.Role
	err := r.db.Where("id = $1", id).First(&role).Error
	return role, err
}

func (r *roleRepository) GetRoleByName(roleName string) (model.Role, error) {
	var role model.Role
	err := r.db.Where("role_name LIKE $1", "%"+roleName+"%").Find(&role).Error
	return role, err
}

func (r *roleRepository) List() ([]model.Role, error) {
	var role []model.Role
	result := r.db.Find(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}

func (r *roleRepository) Update(payload model.Role) error {
	err := r.db.Model(&payload).Updates(payload).Error
	return err
}

func (r *roleRepository) Delete(id string) error {
	role := model.Role{}
	err := r.db.Where("id = $1", id).Delete(&role).Error
	return err
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}
