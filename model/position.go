package model

type Position struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsManager bool   `json:"is_manager"`
}
