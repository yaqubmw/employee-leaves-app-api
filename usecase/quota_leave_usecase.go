package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type QuotaLeaveUseCase interface {
	RegisterNewQuotaLeave(payload model.QuotaLeave) error
	FindAllQuotaLeave() ([]model.QuotaLeave, error)
	FindByIdQuotaLeave(id string) (model.QuotaLeave, error)
	UpdateQuotaLeave(payload model.QuotaLeave) error
	DeleteQuotaLeave(id string) error
}

type quotaLeaveUseCase struct {
	repo repository.QuotaLeaveRepository
}


func (q *quotaLeaveUseCase) RegisterNewQuotaLeave(payload model.QuotaLeave) error {
	err := q.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new quota leave: %v", err)
	}
	return nil
}


func (q *quotaLeaveUseCase) FindAllQuotaLeave() ([]model.QuotaLeave, error) {
	return q.repo.List()
}


func (q *quotaLeaveUseCase) FindByIdQuotaLeave(id string) (model.QuotaLeave, error) {
	return q.repo.Get(id)
}



func (q *quotaLeaveUseCase) UpdateQuotaLeave(payload model.QuotaLeave) error {
	err := q.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update quota leave: %v", err)
	}
	return nil
}

func (q *quotaLeaveUseCase) DeleteQuotaLeave(id string) error {
	quotaLeave, err := q.FindByIdQuotaLeave(id)
	if err != nil {
		return fmt.Errorf("data with ID %s not found", id)
	}

	err = q.repo.Delete(quotaLeave.ID)
	if err != nil {
		return fmt.Errorf("failed to delete quota leave: %v", err)
	}
	return nil
}

func NewQuotaLeaveUseCase(repo repository.QuotaLeaveRepository) QuotaLeaveUseCase {
	return &quotaLeaveUseCase{
		repo: repo,
	}
}
