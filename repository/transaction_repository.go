package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload model.TransactionLeave) error
	Get(id string) (dto.TransactionResponseDto, error)
	List() ([]dto.TransactionResponseDto, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func (t *transactionRepository) Create(payload model.TransactionLeave) error {
	err := t.db.Create(&payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) Get(id string) (dto.TransactionResponseDto, error) {
	var transactionResponseDto dto.TransactionResponseDto

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Where("id = ?", id).
		First(&transactionResponseDto).Error
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return transactionResponseDto, nil
}

func (t *transactionRepository) List() ([]dto.TransactionResponseDto, error) {
	var transactionList []dto.TransactionResponseDto

	err := t.db.Preload("Employee").Preload("LeaveType").Preload("StatusLeave").
		Find(&transactionList).Error
	if err != nil {
		return nil, err
	}

	return transactionList, nil
}

func NewTransactionLeaveRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
