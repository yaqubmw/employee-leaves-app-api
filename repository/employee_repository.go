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
	Update(payload model.Employee) error
	Delete(id string) error
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
	var empl model.Employee
	err := e.db.Where("id = $1", id).First(&empl).Error
	return empl, err
}

func (e *emplRepository) GetByName(name string) (model.Employee, error) {
	var empl model.Employee
	err := e.db.Where("name ILIKE $1", "%"+name+"%").First(&empl).Error
	return empl, err
}

func (e *emplRepository) Update(payload model.Employee) error {
	return e.db.Model(&payload).Updates(model.Employee{Name: payload.Name, PhoneNumber: payload.PhoneNumber, Email: payload.Email, Address: payload.Address}).Error
}

func (e *emplRepository) Delete(id string) error {
	return e.db.Where("id = ?", id).Delete(&model.Employee{}).Error
}

func NewEmplRepository(db *gorm.DB) EmplRepository {
	return &emplRepository{db: db}
}
