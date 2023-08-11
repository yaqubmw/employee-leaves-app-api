package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
	"fmt"
)

type EmployeeUseCase interface {
	RegisterNewEmpl(payload model.Employee) error
	FindAllEmpl(requesPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error)
	FindByIdEmpl(id string) (model.Employee, error)
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

func (p *employeeUseCase) FindAllEmpl(requesPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	return p.repo.Paging(requesPaging)
}

func (e *employeeUseCase) FindByIdEmpl(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
