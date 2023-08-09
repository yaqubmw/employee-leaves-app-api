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

func (e *emplUseCase) FindAllEmpl() ([]model.Employee, error) {
	return e.repo.List()
}

func (e *emplUseCase) FindByIdEmpl(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

func NewEmplUseCase(repo repository.EmplRepository) EmplUseCase {
	return &emplUseCase{repo: repo}
}
