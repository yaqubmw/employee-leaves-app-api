package model

import "time"

type TransactionLeave struct {
	ID             int
	EmployeeID     int
	LeaveTypeID    int
	StatusLeaveID  int
	DateStart      time.Time
	DateEnd        time.Time
	TypeOfDay      string
	Reason         string
	SubmissionDate time.Time
}
