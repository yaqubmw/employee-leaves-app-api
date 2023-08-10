package repository

import (
	"employeeleave/model"

	"gorm.io/gorm"
)

type QuotaLeaveRepository interface {
	BaseRepository[model.QuotaLeave]
}

type quotaLeaveRepository struct {
	db *gorm.DB
}

func (q *quotaLeaveRepository) Create(payload model.QuotaLeave) error {
	return q.db.Create(&payload).Error
}

func (q *quotaLeaveRepository) Get(id string) (model.QuotaLeave, error) {
	var quotaLeave model.QuotaLeave
	err := q.db.Where("id = $1", id).First(&quotaLeave).Error

	return quotaLeave, err
}

func (q *quotaLeaveRepository) List() ([]model.QuotaLeave, error) {
	var quotaLeaves []model.QuotaLeave
	err := q.db.Find(&quotaLeaves).Error

	return quotaLeaves, err
}

func (q *quotaLeaveRepository) Update(payload model.QuotaLeave) error {
	err := q.db.Model(&payload).Updates(payload).Error

	return err
}

func (q *quotaLeaveRepository) Delete(id string) error {
	quotaLeave := model.QuotaLeave{}
	err := q.db.Where("id = $1", id).Delete(&quotaLeave).Error

	return err
}

func NewQuotaLeaveRepository(db *gorm.DB) QuotaLeaveRepository {
	return &quotaLeaveRepository{
		db: db,
	}
}
