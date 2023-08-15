package model

type Role struct {
	Id       string `json:"id"`
	RoleName string `json:"role_name"`
}

func (Role) TableName() string {
	return "role"
}
