package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
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
