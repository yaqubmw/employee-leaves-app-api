package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
	"fmt"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(payload model.Employee) error
	FindAllEmployee(requesPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error)
	FindByIdEmployee(id string) (model.Employee, error)
	UpdateEmployee(payload model.Employee) error
	DeleteEmployee(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func (e *employeeUseCase) DeleteEmployee(id string) error {
	employee, err := e.FindByIdEmployee(id)
	if err != nil {
		return fmt.Errorf("employee with ID %s not found", id)
	}

	err = e.repo.Delete(employee.ID)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %v", err.Error())
	}
	return nil
}

func (e *employeeUseCase) FindAllEmployee(requesPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	return e.repo.Paging(requesPaging)
}

func (e *employeeUseCase) FindByIdEmployee(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

func (e *employeeUseCase) RegisterNewEmployee(payload model.Employee) error {
	if payload.Name == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("name, phone number are required fields")
	}
	// employee, _ := e.repo.GetPhoneNumber(payload.PhoneNumber)
	// if employee.PhoneNumber == payload.PhoneNumber {
	// 	return fmt.Errorf("employee with phone number %s already exists", payload.PhoneNumber)
	// }
	err := e.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create employee: %v", err.Error())
	}
	return nil
}

func (e *employeeUseCase) UpdateEmployee(payload model.Employee) error {
	if payload.Name == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("name, phone number are required fields")
	}
	// employee, _ := e.repo.GetPhoneNumber(payload.PhoneNumber)
	// if employee.PhoneNumber == payload.PhoneNumber && employee.Id != payload.Id {
	// 	return fmt.Errorf("employee with phone number %s already exists", payload.PhoneNumber)
	// }
	err := e.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update employee: %v", err.Error())
	}
	return nil
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
