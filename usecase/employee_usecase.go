package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type EmplUseCase interface {
	RegisterNewEmpl(payload model.Employee) error
	FindAllEmpl() ([]model.Employee, error)
	FindByIdEmpl(id string) (model.Employee, error)
	UpdateEmpl(payload model.Employee) error
	DeleteEmpl(id string) error
}

type emplUseCase struct {
	repo repository.EmplRepository
}

// RegisterNewUom implements UomUseCase.
func (e *emplUseCase) RegisterNewEmpl(payload model.Employee) error {
	//pengecekan nama tidak boleh kosong
	if payload.Name == "" {
		return fmt.Errorf("name required fields")
	}

	//pengecekan nama tidak boleh sama
	isExistEmpl, _ := e.repo.GetByName(payload.Name)
	if isExistEmpl.Name == payload.Name {
		return fmt.Errorf("employee with name %s exists", payload.Name)
	}

	err := e.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new employee: %v", err)
	}
	return nil
}

// FindAllUom implements UomUseCase.
func (e *emplUseCase) FindAllEmpl() ([]model.Employee, error) {
	return e.repo.List()
}

// FindByIdUom implements UomUseCase.
func (e *emplUseCase) FindByIdEmpl(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

// DeleteUom implements UomUseCase.
func (e *emplUseCase) DeleteEmpl(id string) error {
	// cek id nya ada atau tidak
	uom, err := e.FindByIdEmpl(id)
	if err != nil {
		return fmt.Errorf("data with ID %s not found", id)
	}

	err = e.repo.Delete(uom.ID)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %v", err)
	}
	return nil
}

// UpdateUom implements UomUseCase.
func (e *emplUseCase) UpdateEmpl(payload model.Employee) error {
	if payload.Name == "" {
		return fmt.Errorf("name is required field")
	}

	isExistEmpl, _ := e.repo.GetByName(payload.Name)
	if isExistEmpl.Name == payload.Name && isExistEmpl.ID != payload.ID {
		return fmt.Errorf("employee with name %s exists", payload.Name)
	}

	err := e.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update employee: %v", err)
	}

	return nil
}

func NewEmplUseCase(repo repository.EmplRepository) EmplUseCase {
	return &emplUseCase{repo: repo}
}
