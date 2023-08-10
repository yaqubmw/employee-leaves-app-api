package model

import "time"

type LeaveApproval struct {
	ID               uint
	LeaveApplication LeaveApplication
	Supervisor       Employee
	Status           string
	ApprovalDate     time.Time
}
