package model

import "time"

type TransactionLeave struct {
	ID             string
	EmployeeID     string
	LeaveTypeID    string
	StatusLeaveID  string
	DateStart      time.Time
	DateEnd        time.Time
	TypeOfDay      string
	Reason         string
	SubmissionDate time.Time
}
