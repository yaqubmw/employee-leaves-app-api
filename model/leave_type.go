package model

type LeaveType struct {
	ID            string `json:"id"`
	LeaveTypeName string `json:"leave_type_name"`
	QuotaLeave    int    `json:"quota_leave"`
}

func (LeaveType) TableName() string {
	return "leave_type"
}
