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
	UpdateAnnualLeave(id string, availableDays int) error
	UpdateMaternityLeave(id string, availableDays int) error
	UpdateMarriageLeave(id string, availableDays int) error
	UpdateMenstrualLeave(id string, availableDays int) error
	PaternityLeave(id string, availableDays int) error
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
	if err := e.db.Where("id = ?", id).First(&employee).Error; err != nil {
		return employee, err
	}
	return employee, nil
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

func (e *employeeRepository) PaternityLeave(id string, availableDays int) error {
	err := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("paternity_leave", availableDays).Error
	return err
}

func (e *employeeRepository) UpdateAnnualLeave(id string, availableDays int) error {
	err := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("annual_leave", availableDays).Error
	return err
}

func (e *employeeRepository) UpdateMarriageLeave(id string, availableDays int) error {
	err := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("marriage_leave", availableDays).Error
	return err
}

func (e *employeeRepository) UpdateMaternityLeave(id string, availableDays int) error {
	err := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("maternity_leave", availableDays).Error
	return err
}

func (e *employeeRepository) UpdateMenstrualLeave(id string, availableDays int) error {
	err := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("menstrual_leave", availableDays).Error
	return err
}

func NewEmplRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
