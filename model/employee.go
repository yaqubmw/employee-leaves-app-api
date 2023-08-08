package model

type Employee struct {
	ID          string `json:"id"`
	PositionID  string `json:"position_id"`
	ManagerID   string `json:"manager_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
