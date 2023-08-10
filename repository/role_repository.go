package repository

import (
	"employeeleave/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[model.Role]
	GetByName(roleName string) (model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) Create(payload model.Role) error {
	return r.db.Create(&payload).Error
}

func (r *roleRepository) Get(id string) (model.Role, error) {
	var role model.Role
	err := r.db.First(&role, id).Error
	return role, err
}

func (r *roleRepository) GetByName(roleName string) (model.Role, error) {
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
	var role model.Role
	result := r.db.Where("id = ?", id).Delete(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}
