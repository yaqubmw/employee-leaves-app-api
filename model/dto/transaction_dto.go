package dto

import (
	"employeeleave/model"
	"time"
)

type TransactionResponseDto struct {
	ID             string            `json:"id"`
	DateStart      string            `json:"dateStart"`
	DateEnd        string            `json:"dateEnd"`
	DayType        string            `json:"dayType"`
	Reason         string            `json:"reason"`
	SubmissionDate time.Time         `json:"submissionDate"`
	Employee       model.Employee    `json:"employee" gorm:"foreignKey:EmployeeID"`
	EmployeeID     string            `json:"-"`
	LeaveType      model.LeaveType   `json:"leaveType" gorm:"foreignKey:LeaveTypeID"`
	LeaveTypeID    string            `json:"-"`
	StatusLeave    model.StatusLeave `json:"statusLeave" gorm:"foreignKey:StatusLeaveID"`
	StatusLeaveID  string            `json:"-"`
}
