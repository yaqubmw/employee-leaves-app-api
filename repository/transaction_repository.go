package repository

import (
	"database/sql"
	"employeeleave/model"
	"employeeleave/model/dto"
)

type TransactionRepository interface {
	Create(payload model.TransactionLeave) error
	Get(id string) (dto.TransactionResponseDto, error)
	List() ([]dto.TransactionResponseDto, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) Create(payload model.TransactionLeave) error {
	tx, err := t.db.Begin()

	if err != nil {
		return err
	}
	// insert transaction
	_, err = tx.Exec("INSERT INTO transaction_leave (id, date_start, date_end, type_of_day, reason, submission_date) VALUES ($1, $2, $3, $4, $5, $6)",
		payload.ID, payload.DateStart, payload.DateEnd, payload.TypeOfDay, payload.Reason, payload.SubmissionDate)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (t *transactionRepository) Get(id string) (dto.TransactionResponseDto, error) {
	panic("not implemented")
}

func (t *transactionRepository) List() ([]dto.TransactionResponseDto, error) {
	panic("not implemented")
}

func NewTransactionLeaveRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
