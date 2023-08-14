package model

type StatusLeave struct {
	ID              string
	StatusLeaveName string
}

func (StatusLeave) TableName() string {
	return "status_leave"
}
