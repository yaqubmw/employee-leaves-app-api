package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/utils/common"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload model.TransactionLeave) error
	GetByID(id string) (model.TransactionLeave, error)
	GetByEmployeeID(employeeID string) ([]model.TransactionLeave, error)
	UpdateStatus(transactionID, statusID string) error
	GetByIdTxNonDto(id string) (model.TransactionLeave, error)
	BaseRepositoryPaging[dto.TransactionResponseDto]
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
func (t *transactionRepository) GetByID(id string) (model.TransactionLeave, error) {
	var transactionResponse model.TransactionLeave

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Where("id = ?", id).
		First(&transactionResponse).Error
	if err != nil {
		return model.TransactionLeave{}, err
	}

	return transactionResponse, nil
}

// mendapatkan data transaksi cuti berdasarkan ID karyawan
func (t *transactionRepository) GetByEmployeeID(employeeID string) ([]model.TransactionLeave, error) {
	var transactions []model.TransactionLeave

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Where("employee_id = ?", employeeID).
		Find(&transactions).Error
	if err != nil {
		return []model.TransactionLeave{}, err
	}

	return transactions, nil
}

func (t *transactionRepository) Paging(requestPaging dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error) {
	var employees []dto.TransactionResponseDto
	var totalRows int64

	// Dapatkan parameter paginasi
	paginationQuery := common.GetPaginationParams(requestPaging)
	// Ambil data karyawan berdasarkan paginasi menggunakan GORM
	result := t.db.Model(&dto.TransactionResponseDto{}).Count(&totalRows)
	if result.Error != nil {
		return nil, dto.Paging{}, result.Error
	}

	// Hitung total baris data karyawan
	query := t.db.Model(&dto.TransactionResponseDto{}).Limit(paginationQuery.Take).Offset(paginationQuery.Skip)
	result = query.Find(&employees)
	if result.Error != nil {
		return nil, dto.Paging{}, result.Error
	}

	// Kembalikan data karyawan dan informasi paginasi, konversi totalRows menjadi int
	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, int(totalRows)), nil
}

// mengubah status transaksi cuti
func (t *transactionRepository) UpdateStatus(transactionID, statusID string) error {
	err := t.db.Model(&model.TransactionLeave{}).
		Where("id = $1", transactionID).
		Update("status_leave_id", statusID).
		Error
	return err
}

// GetById dari model TransactionLeave (bukan dto)
func (t *transactionRepository) GetByIdTxNonDto(id string) (model.TransactionLeave, error) {
	var txLeave model.TransactionLeave
	err := t.db.Where("id = $1", id).First(&txLeave).Error

	return txLeave, err
}

func NewTransactionLeaveRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
