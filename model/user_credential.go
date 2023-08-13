package model

type UserCredential struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   string `json:"roleId"`
	Role     Role   `gorm:"foreignkey:RoleId"`
}

func (UserCredential) TableName() string {
	return "user_credential"
}
