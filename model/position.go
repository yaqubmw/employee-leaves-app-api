package model

type Position struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Position) TableName() string {
	return "position"
}
