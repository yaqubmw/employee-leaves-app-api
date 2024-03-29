package model

import (
	"time"
)

type TransactionLeave struct {
	ID             string       `json:"id"`
	EmployeeID     string       `json:"employee_id"`
	LeaveTypeID    string       `json:"leave_type_id"`
	StatusLeaveID  string       `json:"status_leave_id"`
	Employee       Employee     `gorm:"foreignkey:EmployeeID"`
	LeaveType      LeaveType    `gorm:"foreignkey:LeaveTypeID"`
	StatusLeave    StatusLeave  `gorm:"foreignkey:StatusLeaveID"`
	DateStart      time.Time    `json:"date_start"`
	DateEnd        time.Time    `json:"date_end"`
	Reason         string       `json:"reason"`
	SubmissionDate time.Time    `json:"submissionDate"`
	AmountLeave    int          `json:"amount_leave"`
	TypeOfDay      string       `json:"type_of_day"`
	HistoryLeaves  HistoryLeave `gorm:"foreignkey:TransactionLeaveId;references:ID"`
}

// nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (TransactionLeave) TableName() string {
	return "transaction_leave"
}
