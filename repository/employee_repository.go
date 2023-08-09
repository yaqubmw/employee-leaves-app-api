package repository

import (
	"employeeleave/model"

	"gorm.io/gorm"
)

type EmplRepository interface {
	Create(payload model.Employee) error
	List() ([]model.Employee, error)
	Get(id string) (model.Employee, error)
	GetByName(name string) (model.Employee, error)
}

type emplRepository struct {
	db *gorm.DB
}

func (e *emplRepository) Create(payload model.Employee) error {
	return e.db.Create(&payload).Error
}

func (e *emplRepository) List() ([]model.Employee, error) {
	var empls []model.Employee
	err := e.db.Find(&empls).Error
	return empls, err
}

func (e *emplRepository) Get(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.First(&employee, id).Error
	return employee, err
}

func (e *emplRepository) GetByName(name string) (model.Employee, error) {
	var employees model.Employee
	err := e.db.Where("name LIKE ?", "%"+name+"%").Find(&employees).Error
	return employees, err
}

func NewEmplRepository(db *gorm.DB) EmplRepository {
	return &emplRepository{db: db}
}
