package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/utils/common"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(payload model.HistoryLeave) error
	GetHistoryById(id string) (model.HistoryLeave, error)
	BaseRepositoryPaging[model.HistoryLeave]
}

type historyRepository struct {
	db *gorm.DB
}

func (h *historyRepository) Create(payload model.HistoryLeave) error {
	return h.db.Create(&payload).Error
}

func (h *historyRepository) GetHistoryById(id string) (model.HistoryLeave, error) {
	var history model.HistoryLeave
	err := h.db.Where("id = $1", id).First(&history).Error

	return history, err
}

func (h *historyRepository) Paging(requestPaging dto.PaginationParam) ([]model.HistoryLeave, dto.Paging, error) {
	var histories []model.HistoryLeave
	var totalRows int64

	pagination := common.GetPaginationParams(requestPaging)
	result := h.db.Model(&model.HistoryLeave{}).Count(&totalRows)
	if result.Error != nil {
		return nil, dto.Paging{}, result.Error
	}

	query := h.db.Model(&model.HistoryLeave{}).Limit(pagination.Take).Offset(pagination.Skip)
	result = query.Find(&histories)
	if result.Error != nil {
		return nil, dto.Paging{}, result.Error
	}

	return histories, common.Paginate(pagination.Page, pagination.Take, int(totalRows)), nil
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}