package model

type Position struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsManager bool   `json:"isManager"`
}

func (Position) TableName() string {
	return "position"
}
