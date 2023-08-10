package repository

import (
	"employeeleave/model"
	"fmt"

	"gorm.io/gorm"
)

type LeaveTypeRepository interface {
	BaseRepository[model.LeaveType]
	GetByName(name string) (model.LeaveType, error)
}

type leaveTypeRepository struct {
	db *gorm.DB
}

func (lt *leaveTypeRepository) Create(payload model.LeaveType) error {
	result := lt.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("leave type created successfully")
	return nil
}

func (lt *leaveTypeRepository) List() ([]model.LeaveType, error) {
	var leavetypes []model.LeaveType
	result := lt.db.Find(&leavetypes)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("leave type retrieve all successfully")
	return leavetypes, nil
}

func (lt *leaveTypeRepository) Get(id string) (model.LeaveType, error) {
	var leave_type model.LeaveType
	result := lt.db.First(&leave_type, id)
	if result.Error != nil {
		return model.LeaveType{}, result.Error
	}

	return leave_type, nil
}

func (lt *leaveTypeRepository) GetByName(leave_type_name string) (model.LeaveType, error) {
	var leave_type model.LeaveType
	result := lt.db.Where("leave_type_name ILIKE ?", "%"+leave_type_name+"%").First(&leave_type)
	if result.Error != nil {
		return model.LeaveType{}, result.Error
	}

	return leave_type, nil
}

func (lt *leaveTypeRepository) Update(payload model.LeaveType) error {
	result := lt.db.Model(&payload).Updates(model.LeaveType{LeaveTypeName: payload.LeaveTypeName, QuotaLeave: payload.QuotaLeave})
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Successfully Updated")
	return nil
}

func (lt *leaveTypeRepository) Delete(id string) error {
	result := lt.db.Delete(&model.LeaveType{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewLeaveTypeRepository(db *gorm.DB) LeaveTypeRepository {
	return &leaveTypeRepository{db: db}
}
