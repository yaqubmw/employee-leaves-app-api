package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
)

type HistoryUseCase interface {
	RegisterNewHistory(payload model.HistoryLeave) error
	FindAllHistory() ([]model.HistoryLeave, error)
	FindByHistoryId(id string) (model.HistoryLeave, error)
}

type historyUseCase struct {
	repo repository.HistoryRepository
}

func (h *historyUseCase) FindAllHistory() ([]model.HistoryLeave, error) {
	return h.repo.List()
}

func (h *historyUseCase) FindByHistoryId(id string) (model.HistoryLeave, error) {
	return h.repo.GetHistory(id)
}

func (h *historyUseCase) RegisterNewHistory(payload model.HistoryLeave) error {
	return h.repo.Create(payload)
}

func NewHistoryUseCase(repo repository.HistoryRepository) HistoryUseCase {
	return &historyUseCase{repo: repo}
}
