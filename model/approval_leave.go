package model

import "time"

type ApprovalLeave struct {
	ID            string
	TransactionID string
	PositionID    string
	DateApproval  time.Time
}
