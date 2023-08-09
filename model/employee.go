package model

type Employee struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

// Pastikan bahwa nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (Employee) TableName() string {
	return "employee"
}
