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

func (suite *EmployeeRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

var employeDumy = []model.Employee{
	{
		ID:               "1",
		Name:             "imron",
		PositionID:       "1",
		UserCredentialID: "1",
		PhoneNumber:      "1234567890",
		Email:            "im@example.com",
		Address:          "jakarta St",
		AnnualLeave:      0,
		MaternityLeave:   0,
		MarriageLeave:    0,
		MenstrualLeave:   0,
		PaternityLeave:   0,
	},
	{
		ID:               "2",
		Name:             "imam ",
		PositionID:       "2",
		UserCredentialID: "2",
		PhoneNumber:      "12314",
		Email:            "imam@example.com",
		Address:          "jakarta",
		AnnualLeave:      0,
		MaternityLeave:   0,
		MarriageLeave:    0,
		MenstrualLeave:   0,
		PaternityLeave:   0,
	},
}

func (suite *EmployeeRepositorySuite) TestCreate() {
	payload := employeDumy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"employee\" (.+)").WithArgs(
		payload.ID,
		payload.Name,
		payload.PositionID,
		payload.UserCredentialID,
		payload.PhoneNumber,
		payload.Email,
		payload.Address,
		payload.AnnualLeave,
		payload.MaternityLeave,
		payload.MarriageLeave,
		payload.MenstrualLeave,
		payload.PaternityLeave,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)

}

func (suite *EmployeeRepositorySuite) TestList() {
	expectedEmployee := employeDumy

	rows := sqlmock.NewRows([]string{"id", "name", "position_id", "user_credential_id", "phone_number", "email", "address", "annual_leave", "maternity_leave", "marriage_leave", "menstrual_leave", "paternity_leave"})

	for _, emp := range expectedEmployee {
		rows.AddRow(emp.ID, emp.Name, emp.PositionID, emp.UserCredentialID, emp.PhoneNumber, emp.Email, emp.Address, emp.AnnualLeave, emp.MaternityLeave, emp.MarriageLeave, emp.MenstrualLeave, emp.PaternityLeave)
	}

	expectedQuery := `SELECT \* FROM "employee"`

	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

	result, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedEmployee, result)
}

func (suite *EmployeeRepositorySuite) TestGetById() {
	employeeID := "1"
	expectedEmployee := model.Employee{
		ID:   "1",
		Name: "John Doe",
	}
	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(expectedEmployee.ID, expectedEmployee.Name)
	expectedQuery := `SELECT \* FROM "employee" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(employeeID).WillReturnRows(rows)

	result, err := suite.repo.Get(employeeID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedEmployee, result)
}

func (suite *EmployeeRepositorySuite) TestGetByName() {
	employeeName := "imron"
	expectedEmployee := model.Employee{
		ID:   "1",
		Name: "imron",
	}

	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(expectedEmployee.ID, expectedEmployee.Name)
	expectedQuery := `SELECT \* FROM "employee" WHERE name LIKE \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs("%" + employeeName + "%").WillReturnRows(rows)

	result, err := suite.repo.GetByName(employeeName)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedEmployee, result)
}

func (suite *EmployeeRepositorySuite) TestUpdate() {
	// Define the updated employee data
	updatedEmployee := employeDumy[0]
	updatedEmployee.Name = "imron"
	updatedEmployee.Email = "imron@example.com"

	// Set up the mock expectation for the UPDATE query
	expectedQuery := `UPDATE "employee" SET "id"=\$1,"name"=\$2,"position_id"=\$3,"user_credential_id"=\$4,"phone_number"=\$5,"email"=\$6,"address"=\$7 WHERE "id" = \$8`
	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(
		updatedEmployee.ID,
		updatedEmployee.Name,
		updatedEmployee.PositionID,
		updatedEmployee.UserCredentialID,
		updatedEmployee.PhoneNumber,
		updatedEmployee.Email,
		updatedEmployee.Address,
		updatedEmployee.ID,
	).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	// Call the Update method
	err := suite.repo.Update(updatedEmployee)
	assert.Nil(suite.T(), err)
}

func (suite *EmployeeRepositorySuite) TestUpdatePaternityLeave() {
	employeeID := employeDumy[0].ID
	availableDays := 0

	expectedQuery := `UPDATE "employee" SET "paternity_leave"=\$1 WHERE id = \$2`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(availableDays, employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.UpdatePaternityLeave(employeeID, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeRepositorySuite) TestUpdateUpdateAnnualLeave() {
	employeeID := employeDumy[0].ID
	availableDays := 0

	expectedQuery := `UPDATE "employee" SET "annual_leave"=\$1 WHERE id = \$2`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(availableDays, employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.UpdateAnnualLeave(employeeID, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeRepositorySuite) TestUpdateMarriageLeave() {
	employeeID := employeDumy[0].ID
	availableDays := 0

	expectedQuery := `UPDATE "employee" SET "marriage_leave"=\$1 WHERE id = \$2`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(availableDays, employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.UpdateMarriageLeave(employeeID, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeRepositorySuite) TestUpdateMaternityLeave() {
	employeeID := employeDumy[0].ID
	availableDays := 0

	expectedQuery := `UPDATE "employee" SET "maternity_leave"=\$1 WHERE id = \$2`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(availableDays, employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.UpdateMaternityLeave(employeeID, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeRepositorySuite) TestUpdateMenstrualLeave() {
	employeeID := employeDumy[0].ID
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
