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
	Employee       model.Employee    `json:"employee"`
	LeaveType      model.LeaveType   `json:"leaveType"`
	StatusLeave    model.StatusLeave `json:"statusLeave"`
}

func (TransactionResponseDto) TableName() string {
	return "transaction_response_dto"
}