package model

type UserCredential struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	RoleId   string `json:"roleId"`
	IsActive bool   `json:"is_active"`
	Role     Role   `gorm:"foreignkey:RoleId"`
}

func (UserCredential) TableName() string {
	return "user_credential"
}
