package model

import "time"

type HistoryLeave struct {
	Id            string `json:"id"`
	EmployeeId     string `json:"employeeId"`
	TransactionId     string `json:"transactionId"`
	// Employee    	Employee `json:"employee"`
	// Transaction 	TransactionLeave `json:"transactionLeave"`
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
	LeaveDuration string `json:"leaveDuration"`
	StatusLeave   string `json:"statusDuration"`
}

func (HistoryLeave) TableName() string {
	return "history_leave"
}