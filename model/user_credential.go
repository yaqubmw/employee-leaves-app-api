package model

type UserCredential struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   string `json:"roleId"`
}

func (UserCredential) TableName() string {
	return "user_credential"
}
