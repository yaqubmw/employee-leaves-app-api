package model

import "time"

type HistoryLeave struct {
	ID            int
	EmployeeID    int
	TransactionID int
	DateStart     time.Time
	DateEnd       time.Time
	LeaveDuration string
	StatusLeave   string
}
