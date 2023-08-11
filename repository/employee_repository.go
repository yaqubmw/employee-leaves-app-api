package repository

import (
	"employeeleave/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(payload model.Employee) error
	List() ([]model.Employee, error)
	Get(id string) (model.Employee, error)
	GetByName(name string) (model.Employee, error)
	Update(employee model.Employee) error
	UpdateAvailableDays(id string, availableDays int) error
}

type employeeRepository struct {
	db *gorm.DB
}

func (e *employeeRepository) Create(payload model.Employee) error {
	return e.db.Create(&payload).Error
}

func (e *employeeRepository) List() ([]model.Employee, error) {
	var empls []model.Employee
	err := e.db.Find(&empls).Error
	return empls, err
}

func (e *employeeRepository) Get(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.First(&employee, id).Error
	return employee, err
}

func (e *employeeRepository) GetByName(name string) (model.Employee, error) {
	var employees model.Employee
	err := e.db.Where("name LIKE $1", "%"+name+"%").Find(&employees).Error
	return employees, err
}

func (e *employeeRepository) Update(employee model.Employee) error {
	err := e.db.Model(&employee).Updates(employee).Error
	return err
}

func (e *employeeRepository) UpdateAvailableDays(id string, availableDays int) error {
	err := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("available_leave_days", availableDays).Error
	return err
}

func NewEmplRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
