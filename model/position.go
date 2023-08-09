package model

type Position struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	IsManager bool   `json:"isManager"`
}
