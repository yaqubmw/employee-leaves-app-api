package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type RoleUseCase interface {
	RegisterNewRole(payload model.Role) error
	FindAllRole() ([]model.Role, error)
	FindByIdRole(id string) (model.Role, error)
	FindByRolename(roleName string) (model.Role, error)
	UpdateRole(payload model.Role) error
	DeleteRole(id string) error
}

type roleUseCase struct {
	repo repository.RoleRepository
}

func (r *roleUseCase) FindAllRole() ([]model.Role, error) {
	return r.repo.List()
}

func (r *roleUseCase) FindByRolename(roleName string) (model.Role, error) {
	return r.repo.GetRoleByName(roleName)
}

func (r *roleUseCase) FindByIdRole(id string) (model.Role, error) {
	return r.repo.Get(id)
}

func (r *roleUseCase) RegisterNewRole(payload model.Role) error {
	// Pengecekan nama tidak boleh kosong
	if payload.RoleName == "" {
		return fmt.Errorf("required fields")
	}

	// nama tidak boleh sama
	isExistRole, _ := r.repo.GetRoleByName(payload.RoleName)
	if isExistRole.RoleName == payload.RoleName {
		return fmt.Errorf("role with name %s already exist", payload.RoleName)
	}

	err := r.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new Role: %v", err)
	}
	return nil
}

func (r *roleUseCase) UpdateRole(payload model.Role) error {
	if payload.RoleName == "" {
		return fmt.Errorf("role name is required")
	}

	isExistRole, _ := r.repo.GetRoleByName(payload.RoleName)
	if isExistRole.RoleName == payload.RoleName && isExistRole.Id != payload.Id {
		return fmt.Errorf("role with name %s exists", payload.RoleName)
	}

	err := r.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}

	return nil
}

func (r *roleUseCase) DeleteRole(id string) error {
	// cek idnya ada atau tidak
	uom, err := r.FindByIdRole(id)
	if err != nil {
		return fmt.Errorf("data with id %s not found", id)
	}

	err = r.repo.Delete(uom.Id)
	if err != nil {
		return fmt.Errorf("failed to delete uom: %v", err)
	}
	return nil
}

func NewRoleUseCase(repo repository.RoleRepository) RoleUseCase {
	return &roleUseCase{repo: repo}
}
