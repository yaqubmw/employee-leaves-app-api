package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/utils/common"
	"errors"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload model.TransactionLeave) error
	BaseRepositoryPaging[model.TransactionLeave]
	Get(id string) (model.TransactionLeave, error)
	// GetByID(id string) (dto.TransactionResponseDto, error)
	// GetByEmployeeID(employeeID string) ([]dto.TransactionResponseDto, error)
	// GetByName(name string) ([]dto.TransactionResponseDto, error)
	// List(requestPaging dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error)
	// UpdateStatus(transactionID, statusID string) error
	// DeleteByID(id string) error
}

type transactionRepository struct {
	db *gorm.DB
}

// membuat data transaksi cuti baru
func (t *transactionRepository) Create(payload model.TransactionLeave) error {
	tx := t.db.Begin() // Memulai transaksi
	if tx.Error != nil {
		return tx.Error
	}

	// Menyimpan data transaksi dalam transaksi
	if err := tx.Create(&payload).Error; err != nil {
		tx.Rollback() // Rollback jika ada error
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// ]mendapatkan data transaksi cuti berdasarkan ID
// func (t *transactionRepository) GetByID(id string) (dto.TransactionResponseDto, error) {
// 	var transactionResponseDto dto.TransactionResponseDto

// 	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
// 		Where("id = ?", id).
// 		First(&transactionResponseDto).Error
// 	if err != nil {
// 		return dto.TransactionResponseDto{}, err
// 	}

// 	return transactionResponseDto, nil
// }

// mendapatkan data transaksi cuti berdasarkan ID karyawan
// func (t *transactionRepository) GetByEmployeeID(employeeID string) ([]dto.TransactionResponseDto, error) {
// 	var transactions []dto.TransactionResponseDto

// 	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
// 		Where("employee_id = ?", employeeID).
// 		Find(&transactions).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// mendapatkan daftar transaksi cuti berdasarkan nama karyawan
// func (t *transactionRepository) GetByName(name string) ([]dto.TransactionResponseDto, error) {
// 	var transactionList []dto.TransactionResponseDto

// 	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
// 		Joins("JOIN employees ON transaction_leaves.employee_id = employees.id").
// 		Where("employees.name LIKE ?", "%"+name+"%").
// 		Find(&transactionList).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactionList, nil
// }

// mendapatkan daftar semua transaksi cuti
// func (t *transactionRepository) List(requestPaging dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error) {
// 	return nil, dto.Paging{}, nil
// }

// mengubah status transaksi cuti
// func (t *transactionRepository) UpdateStatus(transactionID, statusID string) error {
// 	err := t.db.Model(&model.TransactionLeave{}).
// 		Where("id = ?", transactionID).
// 		Update("status_leave_id", statusID).
// 		Error
// 	return err
// }

// menghapus data transaksi cuti berdasarkan ID
// func (t *transactionRepository) DeleteByID(id string) error {
// 	err := t.db.Where("id = ?", id).Delete(&model.TransactionLeave{}).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (t *transactionRepository) Get(id string) (model.TransactionLeave, error) {
	var employee model.TransactionLeave
	// Menggunakan GORM untuk mencari karyawan berdasarkan ID
	if err := t.db.First(&employee, "id = $1", id).Error; err != nil {
		// Jika data tidak ditemukan, kembalikan error dengan pesan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.TransactionLeave{}, errors.New("Transaction Not Found")
		}
		return model.TransactionLeave{}, err
	}
	return employee, nil
}

func (t *transactionRepository) Paging(requestPaging dto.PaginationParam) ([]model.TransactionLeave, dto.Paging, error) {
	var employees []model.TransactionLeave
	var totalRows int64

	// Dapatkan parameter paginasi
	paginationQuery := common.GetPaginationParams(requestPaging)

	// Ambil data karyawan berdasarkan paginasi menggunakan GORM
	if err := t.db.Model(&model.TransactionLeave{}).
		Select("id, employee_id, leave_type_id, status_leave_id, date_start, date_end, type_of_day, reason, submission_date ").
		Limit(paginationQuery.Take).
		Offset(paginationQuery.Skip).
		Find(&employees).
		Error; err != nil {
		return nil, dto.Paging{}, err
	}

	// Hitung total baris data karyawan
	if err := t.db.Model(&model.TransactionLeave{}).Count(&totalRows).Error; err != nil {
		return nil, dto.Paging{}, err
	}

	// Kembalikan data karyawan dan informasi paginasi, konversi totalRows menjadi int
	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, int(totalRows)), nil
}

func NewTransactionLeaveRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
