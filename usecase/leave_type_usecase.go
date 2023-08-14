package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type LeaveTypeUseCase interface {
	RegisterNewLeaveType(payload model.LeaveType) error
	FindAllLeaveType() ([]model.LeaveType, error)
	FindByIdLeaveType(id string) (model.LeaveType, error)
	UpdateLeaveType(payload model.LeaveType) error
	DeleteLeaveType(id string) error
	GetByName(name string) (model.LeaveType, error)
}

type leaveTypeUseCase struct {
	repo repository.LeaveTypeRepository
}

func (lt *leaveTypeUseCase) RegisterNewLeaveType(payload model.LeaveType) error {
	if payload.LeaveTypeName == "" {
		return fmt.Errorf("name are required fields")
	}
	isExistLeaveType, _ := lt.repo.GetByName(payload.LeaveTypeName)
	if isExistLeaveType.LeaveTypeName == payload.LeaveTypeName {
		return fmt.Errorf("leave type with name %s exists", payload.LeaveTypeName)
	}
	err := lt.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new leave type: %v", err)
	}
	return nil
}

func (lt *leaveTypeUseCase) FindAllLeaveType() ([]model.LeaveType, error) {
	return lt.repo.List()
}

func (lt *leaveTypeUseCase) FindByIdLeaveType(id string) (model.LeaveType, error) {
	return lt.repo.Get(id)
}

func (lt *leaveTypeUseCase) UpdateLeaveType(payload model.LeaveType) error {
	fmt.Println("LeaveTypeUseCase.UpdateLeaveType.payload:", payload)
	return lt.repo.Update(payload)
}

func (lt *leaveTypeUseCase) DeleteLeaveType(id string) error {
	return lt.repo.Delete(id)
}

func (p *leaveTypeUseCase) GetByName(name string) (model.LeaveType, error) {
	return p.repo.GetByName(name)
}

func NewLeaveTypeUseCase(repo repository.LeaveTypeRepository) LeaveTypeUseCase {
	return &leaveTypeUseCase{repo: repo}
}
