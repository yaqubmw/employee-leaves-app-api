package repository

import (
	"employeeleave/model"

	"gorm.io/gorm"
)

type StatusLeaveRepository interface {
	BaseRepository[model.StatusLeave]
	GetByNameStatus(statusLeaveName string) (model.StatusLeave, error)
}

type statusLeaveRepository struct {
	db *gorm.DB
}

func (s *statusLeaveRepository) Create(payload model.StatusLeave) error {
	return s.db.Create(&payload).Error
}

func (s *statusLeaveRepository) Get(id string) (model.StatusLeave, error) {
	var statusLeave model.StatusLeave
	err := s.db.Where("id = ?", id).First(&statusLeave).Error

	return statusLeave, err
}

func (s *statusLeaveRepository) List() ([]model.StatusLeave, error) {
	var statusLeaves []model.StatusLeave
	err := s.db.Find(&statusLeaves).Error

	return statusLeaves, err
}

func (s *statusLeaveRepository) Update(payload model.StatusLeave) error {
	err := s.db.Model(&payload).Updates(payload).Error

	return err
}

func (s *statusLeaveRepository) Delete(id string) error {
	statusLeave := model.StatusLeave{}
	err := s.db.Where("id = $1", id).Delete(&statusLeave).Error

	return err
}

func (r *statusLeaveRepository) GetByNameStatus(statusLeaveName string) (model.StatusLeave, error) {
	var status model.StatusLeave
	err := r.db.Where("status_leave_name LIKE $1", "%"+statusLeaveName+"%").Find(&status).Error
	return status, err
}

func NewStatusLeaveRepository(db *gorm.DB) StatusLeaveRepository {
	return &statusLeaveRepository{
		db: db,
	}
}
