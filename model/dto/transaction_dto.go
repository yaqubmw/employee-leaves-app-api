package dto

import (
	"time"
)

type TransactionResponseDto struct {
	ID             string    `json:"id"`
	EmployeeID     string    `json:"employee_id"`
	LeaveTypeID    string    `json:"leave_type_id"`
	StatusLeaveID  string    `json:"status_leave_id"`
	DateStart      string    `json:"dateStart"`
	DateEnd        string    `json:"dateEnd"`
	DayType        string    `json:"dayType"`
	Reason         string    `json:"reason"`
	SubmissionDate time.Time `json:"submissionDate"`
}

func (TransactionResponseDto) TableName() string {
	return "transaction_leave"
}
