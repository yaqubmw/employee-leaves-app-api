package model

import "time"

type HistoryLeave struct {
	Id                 string    `json:"id"`
	TransactionLeaveId string    `json:"transaction_leave_id"`
	DateEvent          time.Time `json:"date_event"`
}

func (HistoryLeave) TableName() string {
	return "history_leave"
}
