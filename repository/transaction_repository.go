package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload model.TransactionLeave) error
	GetByID(id string) (dto.TransactionResponseDto, error)
	GetByIdTxNonDto(id string) (model.TransactionLeave, error)
	GetByEmployeeID(employeeID string) ([]dto.TransactionResponseDto, error)
	GetByName(name string) ([]dto.TransactionResponseDto, error)
	List() ([]dto.TransactionResponseDto, error)
	UpdateStatus(transactionID, statusID string) error
	DeleteByID(id string) error
}

type transactionRepository struct {
	db *gorm.DB
}

// GetByIdTxNonDto implements TransactionRepository.
func (t *transactionRepository) GetByIdTxNonDto(id string) (model.TransactionLeave, error) {
	var txLeave model.TransactionLeave
	err := t.db.Where("id = $1", id).First(&txLeave).Error

	return txLeave, err
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
func (t *transactionRepository) GetByID(id string) (dto.TransactionResponseDto, error) {
	var transactionResponseDto dto.TransactionResponseDto

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Where("id = ?", id).
		First(&transactionResponseDto).Error
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return transactionResponseDto, nil
}

// mendapatkan data transaksi cuti berdasarkan ID karyawan
func (t *transactionRepository) GetByEmployeeID(employeeID string) ([]dto.TransactionResponseDto, error) {
	var transactions []dto.TransactionResponseDto

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Where("employee_id = ?", employeeID).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// mendapatkan daftar transaksi cuti berdasarkan nama karyawan
func (t *transactionRepository) GetByName(name string) ([]dto.TransactionResponseDto, error) {
	var transactionList []dto.TransactionResponseDto

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Joins("JOIN employee ON transaction_leave.employee_id = employee.id").
		Where("employee.name LIKE ?", "%"+name+"%").
		Find(&transactionList).Error
	if err != nil {
		return nil, err
	}

	return transactionList, nil
}

// mendapatkan daftar semua transaksi cuti
func (t *transactionRepository) List() ([]dto.TransactionResponseDto, error) {
	var transactionList []dto.TransactionResponseDto

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Find(&transactionList).Error
	if err != nil {
		return nil, err
	}

	return transactionList, nil
}

// mengubah status transaksi cuti
func (t *transactionRepository) UpdateStatus(transactionID, statusID string) error {
	err := t.db.Model(&model.TransactionLeave{}).
		Where("id = ?", transactionID).
		Update("status_leave_id", statusID).
		Error
	return err
}

// menghapus data transaksi cuti berdasarkan ID
func (t *transactionRepository) DeleteByID(id string) error {
	err := t.db.Where("id = ?", id).Delete(&model.TransactionLeave{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewTransactionLeaveRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
