package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type EmployeeUseCase interface {
	RegisterNewEmpl(payload model.Employee) error
	FindAllEmpl() ([]model.Employee, error)
	FindByIdEmpl(id string) (model.Employee, error)
	UpdateEmpl(payload model.Employee) error
	UpdateAvailableDay(id string, availableDays int) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func (e *employeeUseCase) RegisterNewEmpl(payload model.Employee) error {
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

func (e *employeeUseCase) FindAllEmpl() ([]model.Employee, error) {
	return e.repo.List()
}

func (e *employeeUseCase) FindByIdEmpl(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

func (e *employeeUseCase) UpdateEmpl(payload model.Employee) error {
	err := e.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update employee: %v", err)
	}
	return nil
}

func (e *employeeUseCase) UpdateAvailableDay(id string, availableDays int) error {
	return e.repo.UpdateAvailableDays(id, availableDays)
}

func NewEmplUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
