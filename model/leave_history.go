package model

import "time"

type LeaveHistory struct {
	ID                     uint
	Employee               Employee
	LeaveType              LeaveType
	StartDate              time.Time
	EndDate                time.Time
	FullDayOrHalfDay       string
	LeaveDays              int
	Status                 string
	ApplicationDate        time.Time
	SupervisorApprovalDate time.Time
	HRApprovalDate         time.Time
}
