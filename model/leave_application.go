package model

import "time"

type LeaveApplication struct {
	ID                       uint
	Employee                 Employee
	LeaveType                LeaveType
	Reason                   string
	StartDate                time.Time
	EndDate                  time.Time
	FullDayOrHalfDay         string
	LeaveDays                int
	SupervisorApprovalStatus string
	HRApprovalStatus         string
}
