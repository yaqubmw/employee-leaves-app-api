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
	UpdateAnnualLeave(id string, availableDays int) error
	UpdateMaternityLeave(id string, availableDays int) error
	UpdateMarriageLeave(id string, availableDays int) error
	UpdateMenstrualLeave(id string, availableDays int) error
	PaternityLeave(id string, availableDays int) error
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

func (e *employeeUseCase) PaternityLeave(id string, availableDays int) error {
	return e.repo.PaternityLeave(id, availableDays)
}

func (e *employeeUseCase) UpdateAnnualLeave(id string, availableDays int) error {
	return e.repo.UpdateAnnualLeave(id, availableDays)
}

func (e *employeeUseCase) UpdateMarriageLeave(id string, availableDays int) error {
	return e.repo.UpdateMarriageLeave(id, availableDays)
}

func (e *employeeUseCase) UpdateMaternityLeave(id string, availableDays int) error {
	return e.repo.UpdateMaternityLeave(id, availableDays)
}

func (e *employeeUseCase) UpdateMenstrualLeave(id string, availableDays int) error {
	return e.repo.UpdateMenstrualLeave(id, availableDays)
}

func NewEmplUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
