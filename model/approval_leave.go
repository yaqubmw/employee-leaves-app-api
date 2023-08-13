package model

import "time"

type ApprovalLeave struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	DateApproval  time.Time `json:"date_approval"`
}

func (ApprovalLeave) TableName() string {
	return "approval_leave"
}
