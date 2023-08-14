package repository

import (
	"employeeleave/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EmployeeRepositorySuite struct {
	suite.Suite
	repo    EmployeeRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *EmployeeRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewEmplRepository(gormDB)
}


var employeeDummy = []model.Employee{
	{
		ID:               "1",
		Name:             "Agus",
		PositionID:       "1",
		UserCredentialID: "1",
		PhoneNumber:      "0899776651",
		Email:            "agus@mail.com",
		Address:          "Jakarta",
		AnnualLeave:      12,
		MaternityLeave:   84,
		MarriageLeave:    3,
		MenstrualLeave:   2,
		PaternityLeave:   2,
	},
	{
		ID:               "2",
		Name:             "Septi",
		PositionID:       "1",
		UserCredentialID: "2",
		PhoneNumber:      "0899776666",
		Email:            "septi@mail.com",
		Address:          "Surabaya",
		AnnualLeave:      12,
		MaternityLeave:   84,
		MarriageLeave:    3,
		MenstrualLeave:   2,
		PaternityLeave:   2,
	},
}

func (suite *EmployeeRepositorySuite) TestUpdateMenstrualLeave() {
    employeeID := employeeDummy[0].ID
    availableDays := 0

    expectedQuery := `UPDATE "employee" SET "menstrual_leave"=\$1 WHERE id = \$2`

	suite.mocksql.ExpectBegin()
    suite.mocksql.ExpectExec(expectedQuery).WithArgs(availableDays, employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

    err := suite.repo.UpdateMenstrualLeave(employeeID, availableDays)
    assert.NoError(suite.T(), err)
}

func TestEmployeeRepositorySuite(t *testing.T) {
	suite.Run(t, new(EmployeeRepositorySuite))
}
