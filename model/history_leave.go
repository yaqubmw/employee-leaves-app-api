package model

import "time"

type HistoryLeave struct {
	Id            string `json:"id"`
	Employee    	Employee `json:"employee"`
	Transaction 	TransactionLeave `json:"transactionLeave"`
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
	LeaveDuration string `json:"leaveDuration"`
	StatusLeave   string `json:"statusDuration"`
}
