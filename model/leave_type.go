package model

type LeaveType struct {
	ID            string
	LeaveTypeName string
	QuotaLeave    int
}

func (LeaveType) TableName() string {
	return "leave_type"
}
