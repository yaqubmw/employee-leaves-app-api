package model

type Employee struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	PositionID       string `json:"position_id"`
	UserCredentialID string `json:"user_credential_id"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	Address          string `json:"address"`
	AnnualLeave      int    `json:"annualLeave"`
	MaternityLeave   int    `json:"maternityLeave"`
	MarriageLeave    int    `json:"marriageLeave"`
	MenstrualLeave   int    `json:"menstrualLeave"`
	PaternityLeave   int    `json:"paternityLeave"`
}

// nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (Employee) TableName() string {
	return "employee"
}
