package model

import "time"

type HistoryLeave struct {
	ID            string
	EmployeeID    string
	TransactionID string
	DateStart     time.Time
	DateEnd       time.Time
	LeaveDuration string
	StatusLeave   string
}
