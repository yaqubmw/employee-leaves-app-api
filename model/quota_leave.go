package model

type QuotaLeave struct {
	ID             string `json:"id"`
	RemainingQuota int `json:"remainingQuota"`
}
