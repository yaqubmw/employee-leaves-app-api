package dto

import (
	"employeeleave/model"
	"time"
)

type TransactionResponseDto struct {
	ID             string            `json:"id"`
	DateStart      time.Time         `json:"dateStart"`
	DateEnd        time.Time         `json:"dateEnd"`
	DayType        string            `json:"dayType"`
	Reason         string            `json:"reason"`
	SubmissionDate time.Time         `json:"submissionDate"`
	Employee       model.Employee    `json:"employee"`
	LeaveType      model.LeaveType   `json:"leaveType"`
	StatusLeave    model.StatusLeave `json:"statusLeave"`
}
