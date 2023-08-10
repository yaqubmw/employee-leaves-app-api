package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
)

type HistoryUseCase interface {
	RegisterNewHistory(payload model.HistoryLeave) error
	FindAllHistory(requestPaging dto.PaginationParam) ([]model.HistoryLeave, dto.Paging, error)


	FindHistoryById(id string) (model.HistoryLeave, error)
}

type historyUseCase struct {
	repo repository.HistoryRepository
}

func (h *historyUseCase) FindAllHistory(requestPaging dto.PaginationParam) ([]model.HistoryLeave, dto.Paging, error) {
	return h.repo.Paging(requestPaging)
}

func (h *historyUseCase) FindHistoryById(id string) (model.HistoryLeave, error) {
	return h.repo.GetHistoryById(id)
}

func (h *historyUseCase) RegisterNewHistory(payload model.HistoryLeave) error {
	return h.repo.Create(payload)
}

func NewHistoryUseCase(repo repository.HistoryRepository) HistoryUseCase {
	return &historyUseCase{repo: repo}
}
