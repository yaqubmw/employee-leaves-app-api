package model

import "time"

type HistoryLeave struct {
	Id            string    `json:"id"`
	EmployeeId    string    `json:"employee_id"`
	TransactionId string    `json:"transaction_id"`
	DateStart     time.Time `json:"date_start"`
	DateEnd       time.Time `json:"date_end"`
	LeaveDuration string    `json:"leave_duration"`
	StatusLeave   string    `json:"status_duration"`
}

func (HistoryLeave) TableName() string {
	return "history_leave"
}