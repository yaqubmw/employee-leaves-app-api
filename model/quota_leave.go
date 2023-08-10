package model

type QuotaLeave struct {
	ID             string `json:"id"`
	RemainingQuota int    `json:"remainingQuota"`
}

func (QuotaLeave) TableName() string {
	return "quota_leave"
}
