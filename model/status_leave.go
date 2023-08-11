package model

type StatusLeave struct {
	ID              string `json:"id"`
	StatusLeaveName string `json:"statusLeaveName"`
}

func (StatusLeave) TableName() string {
	return "status_leave"
}
