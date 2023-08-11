package dto

import (
	"time"
)

type HistoryResponseDto struct {
	Id            string    `json:"id"`
	EmployeeId    string    `json:"employeeId"`
	TransactionId string    `json:"transactionId"`
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
	LeaveDuration string    `json:"leaveDuration"`
	StatusLeave   string    `json:"statusLeave"`
}
