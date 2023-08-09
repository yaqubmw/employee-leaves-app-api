package dto

type Employee struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	PhoneNumber string    `gorm:"unique" json:"phoneNumber"`
	Email       string    `gorm:"unique" json:"email"`
	Address     string    `json:"address"`
	PositionID  string    `json:"-"`
	Position    Position  `gorm:"foreignKey:PositionID" json:"position"`
	ManagerID   string    `json:"-"`
	Manager     *Employee `gorm:"foreignKey:ManagerID" json:"manager"`
}

type Position struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
