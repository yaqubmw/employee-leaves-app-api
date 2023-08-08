package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type RoleUseCase interface {
	RegisterNewRole(payload model.Role) error
	FindAllRole() ([]model.Role, error)
	FindByRolename(roleName string) (model.Role, error)
}

type roleUseCase struct {
	repo repository.RoleRepository
}

func (r *roleUseCase) FindAllRole() ([]model.Role, error) {
	return r.repo.List()
}

func (r *roleUseCase) FindByRolename(roleName string) (model.Role, error) {
	return r.repo.GetRole(roleName)
}

func (r *roleUseCase) RegisterNewRole(payload model.Role) error {
	// Pengecekan nama tidak boleh kosong
	if payload.RoleName == "" {
		return fmt.Errorf("required fields")
	}

	// nama tidak boleh sama
	isExistRole, _ := r.repo.GetRole(payload.RoleName)
	if isExistRole.RoleName == payload.RoleName {
		return fmt.Errorf("Role with name %s already exist", payload.RoleName)
	}

	err := r.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new Role: %v", err)
	}
	return nil
}

func NewRoleUseCase(repo repository.RoleRepository) RoleUseCase {
	return &roleUseCase{repo: repo}
}
