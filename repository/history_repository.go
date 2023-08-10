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
	// Paging() ([]model.HistoryLeave, error)
	BaseRepositoryPaging[model.HistoryLeave]
}

type historyRepository struct {
	db *gorm.DB
}

func (h *historyRepository) Create(payload model.HistoryLeave) error {
	return h.db.Create(&payload).Error
	// _, err := h.db.Exec("INSERT INTO history_leave (id, employee_id, transaction_id, date_start, date_end, leave_duration, status_leave) VALUES ($1, $2, $3, $4, $5, $6, $7)", payload.Id, payload.Employee.ID, payload.Transaction.ID, payload.DateStart, payload.DateEnd, payload.LeaveDuration, payload.StatusLeave)
	// if err != nil {
	// 	return err
	// }
	// return nil
}

func (h *historyRepository) GetHistoryById(id string) (model.HistoryLeave, error) {
	// var history model.HistoryLeave
	// err := h.db.First(&history, id).Error
	// return history, err

	var history model.HistoryLeave
	err := h.db.Where("id = $1", id).First(&history).Error

	return history, err

	// var history model.HistoryLeave
	// err := h.db.QueryRow("SELECT id, employee_id, transaction_id, date_start, date_end, leave_duration, status_leave FROM history_leave WHERE id = $1", id).Scan(&history.Id, &history.Employee.ID, &history.Transaction.ID, &history.DateStart, &history.DateEnd, &history.LeaveDuration, &history.StatusLeave)
	// if err != nil {
	// 	return model.HistoryLeave{}, err
	// }
	// return history, nil
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

	// paginationQuery := common.GetPaginationParams(requestPaging)
	// rows, err := h.db.Query("SELECT h.id, e.id, t.id, h.date_start, h.date_end, h.leave_duration, h.status_leave FROM history_leave h INNER JOIN employee e ON e.id = h.employee_id INNER JOIN transaction_leave t ON t.id = h.transaction_id LIMIT $1 OFFSET $2", paginationQuery.Take, paginationQuery.Skip)
	// if err != nil {
	// 	return nil, dto.Paging{}, err
	// }

	// var histories []model.HistoryLeave
	// for rows.Next() {
	// 	var history model.HistoryLeave
	// 	err := rows.Scan(&history.Id, &history.Employee.ID, &history.Transaction.ID, &history.DateStart, &history.DateEnd, &history.LeaveDuration, &history.StatusLeave)
	// 	if err != nil {
	// 		return nil, dto.Paging{}, err
	// 	}
	// 	histories = append(histories, history)
	// }

	// var totalRows int
	// row := h.db.QueryRow("SELECT COUNT(*) FROM history_leave")
	// err = row.Scan(&totalRows)
	// if err != nil {
	// 	return nil, dto.Paging{}, err
	// }
	// return histories, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}


// func (h *historyRepository) List() ([]model.HistoryLeave, error) {
// 	var histories []model.HistoryLeave

// 	rows, err := h.db.Query("SELECT h.id, e.id, t.id, h.date_start, h.date_end, h.leave_duration, h.status_leave FROM history_leave h INNER JOIN employee e ON e.id = h.employee_id INNER JOIN transaction_leave t ON t.id = h.transaction_id")
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var history model.HistoryLeave
// 		err := rows.Scan(&history.Id, &history.Employee.ID, &history.Transaction.ID, &history.DateStart, &history.DateEnd, &history.LeaveDuration, &history.StatusLeave)
// 		if err != nil {
// 			return nil, err
// 		}
// 		histories = append(histories, history)
// 	}
// 	return histories, nil
// }

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}