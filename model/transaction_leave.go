package model

import "time"

type TransactionLeave struct {
	ID             string    `json:"id"`
	EmployeeID     string    `json:"employeeId"`
	LeaveTypeID    string    `json:"leaveTypeId"`
	StatusLeaveID  string    `json:"statusleaveId"`
	DateStart      time.Time `json:"dateStart"`
	DateEnd        time.Time `json:"dateEnd"`
	TypeOfDay      string    `json:"typeOfDay"`
	Reason         string    `json:"reason"`
	SubmissionDate time.Time `json:"submissionDate"`
}

func (TransactionLeave) TableName() string {
	return "transaction_leave"
}
