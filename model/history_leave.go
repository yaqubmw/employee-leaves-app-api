package model

import "time"

type HistoryLeave struct {
	Id            string
	EmployeeId    string
	TransactionId string
	DateStart     time.Time
	DateEnd       time.Time
	LeaveDuration string
	StatusLeave   string
}
