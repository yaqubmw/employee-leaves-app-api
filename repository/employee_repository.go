package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/utils/common"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	BaseRepositoryPaging[model.Employee]
	GetByName(name string) (model.Employee, error)
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
	// Menggunakan GORM untuk mencari karyawan berdasarkan ID
	if err := e.db.First(&employee, "id = $1", id).Error; err != nil {
		// Jika data tidak ditemukan, kembalikan error dengan pesan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Employee{}, errors.New("Employee Not Found")
		}
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) GetByName(name string) (model.Employee, error) {
	var employee model.Employee
	// Menggunakan GORM untuk mencari karyawan berdasarkan nomor telepon
	if err := e.db.Where("name = $1", name).First(&employee).Error; err != nil {
		// Jika data tidak ditemukan, kembalikan error dengan pesan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Employee{}, errors.New("Employee Not Found")
		}
		return model.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) Update(payload model.Employee) error {
	if err := e.db.Model(&payload).Updates(map[string]interface{}{
		"Name":        payload.Name,
		"PhoneNumber": payload.PhoneNumber,
		"Email":       payload.Email,
		"Address":     payload.Address,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) Delete(id string) error {
	var employee model.Employee
	if err := e.db.Where("id = $1", id).Delete(&employee).Error; err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) Paging(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	var employees []model.Employee
	var totalRows int64

	// Dapatkan parameter paginasi
	paginationQuery := common.GetPaginationParams(requestPaging)

	// Ambil data karyawan berdasarkan paginasi menggunakan GORM
	if err := e.db.Model(&model.Employee{}).
		Select("id, name, phone_number, email, address").
		Limit(paginationQuery.Take).
		Offset(paginationQuery.Skip).
		Find(&employees).
		Error; err != nil {
		return nil, dto.Paging{}, err
	}

	// Hitung total baris data karyawan
	if err := e.db.Model(&model.Employee{}).Count(&totalRows).Error; err != nil {
		return nil, dto.Paging{}, err
	}

	// Kembalikan data karyawan dan informasi paginasi, konversi totalRows menjadi int
	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, int(totalRows)), nil
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
