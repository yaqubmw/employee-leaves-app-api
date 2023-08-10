package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type StatusLeaveUseCase interface {
	RegisterNewStatusLeave(payload model.StatusLeave) error
	FindAllStatusLeave() ([]model.StatusLeave, error)
	FindByIdStatusLeave(id string) (model.StatusLeave, error)
	UpdateStatusLeave(payload model.StatusLeave) error
	DeleteStatusLeave(id string) error
}

type statusLeaveUseCase struct {
	repo repository.StatusLeaveRepository
}

func (s *statusLeaveUseCase) RegisterNewStatusLeave(payload model.StatusLeave) error {
	if payload.StatusLeaveName == "" {
		return fmt.Errorf("status-leave-name required field")
	}

	err := s.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new status: %v", err)
	}
	return nil
}

func (s *statusLeaveUseCase) FindAllStatusLeave() ([]model.StatusLeave, error) {
	return s.repo.List()
}

func (s *statusLeaveUseCase) FindByIdStatusLeave(id string) (model.StatusLeave, error) {
	return s.repo.Get(id)
}

func (s *statusLeaveUseCase) UpdateStatusLeave(payload model.StatusLeave) error {
	if payload.StatusLeaveName == "" {
		return fmt.Errorf("status-leave-name required field")
	}

	err := s.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update status: %v", err)
	}
	return nil
}

func (s *statusLeaveUseCase) DeleteStatusLeave(id string) error {
	statusLeave, err := s.FindByIdStatusLeave(id)
	if err != nil {
		return fmt.Errorf("data with ID %s not found", id)
	}

	err = s.repo.Delete(statusLeave.ID)
	if err != nil {
		return fmt.Errorf("failed to delete statusLeave: %v", err)
	}
	return nil
}

func NewStatusLeaveUseCase(repo repository.StatusLeaveRepository) StatusLeaveUseCase {
	return &statusLeaveUseCase{
		repo: repo,
	}
}
