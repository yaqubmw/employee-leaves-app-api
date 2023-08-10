package model

type Role struct {
	Id       string `json:"id"`
	RoleName string `json:"role_name" gorm:"tableName:role"`
}

func (Role) TableName() string {
	return "role"
}
