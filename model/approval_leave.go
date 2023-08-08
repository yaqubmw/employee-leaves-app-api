package model

import "time"

type ApprovalLeave struct {
	ID            int
	TransactionID int
	PositionID    int
	DateApproval  time.Time
}
