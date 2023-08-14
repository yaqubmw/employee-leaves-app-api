package repository

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LeaveTypeRepositorySuite struct {
	suite.Suite
	repo    LeaveTypeRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *LeaveTypeRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewLeaveTypeRepository(gormDB)
}

func (suite *LeaveTypeRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

var dataDummy = []model.LeaveType{
	{
		ID:            "1",
		LeaveTypeName: "Matternity",
		QuotaLeave:    84,
	},
	{
		ID:            "2",
		LeaveTypeName: "Annual",
		QuotaLeave:    12,
	},
}

func (suite *LeaveTypeRepositorySuite) TestCreate_Success() {
	payload := dataDummy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"leave_type\" (.+)").WithArgs(payload.ID, payload.LeaveTypeName, payload.QuotaLeave).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *LeaveTypeRepositorySuite) TestCreate_Failed() {
	payload := dataDummy[0]

	expectedError := fmt.Errorf("failed to create leave type")

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"leave_type\" (.+)").WithArgs(payload.ID, payload.LeaveTypeName, payload.QuotaLeave).WillReturnError(expectedError)
	suite.mocksql.ExpectRollback()

	err := suite.repo.Create(payload)

	assert.EqualError(suite.T(), err, expectedError.Error())
}

func (suite *LeaveTypeRepositorySuite) TestGet_Success() {
	LeaveTypeID := "1"
	expectedLeaveType := dataDummy[0]

	rows := sqlmock.NewRows([]string{"id", "leave_type_name", "quota_leave"})
	rows.AddRow(expectedLeaveType.ID, expectedLeaveType.LeaveTypeName, expectedLeaveType.QuotaLeave)
	expectedQuery := `SELECT \* FROM "leave_type" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(LeaveTypeID).WillReturnRows(rows)

	result, err := suite.repo.Get(LeaveTypeID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLeaveType, result)
}

func (suite *LeaveTypeRepositorySuite) TestGet_Failed() {
	LeaveTypeID := "1"

	expectedError := fmt.Errorf("failed to retrieve leave type")

	expectedQuery := `SELECT \* FROM "leave_type" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(LeaveTypeID).WillReturnError(expectedError)

	result, err := suite.repo.Get(LeaveTypeID)

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.LeaveType{}, result)
}

func (suite *LeaveTypeRepositorySuite) TestList_Success() {
	expectedLeaveTypes := dataDummy

	rows := sqlmock.NewRows([]string{"id", "leave_type_name", "quota_leave"})
	for _, LeaveType := range expectedLeaveTypes {
		rows.AddRow(LeaveType.ID, LeaveType.LeaveTypeName, LeaveType.QuotaLeave)
	}

	expectedQuery := `SELECT \* FROM "leave_type"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

	result, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLeaveTypes, result)
}

func (suite *LeaveTypeRepositorySuite) TestList_Failed() {
	expectedError := fmt.Errorf("failed to retrieve leave types")

	expectedQuery := `SELECT \* FROM "leave_type"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnError(expectedError)

	result, err := suite.repo.List()

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Nil(suite.T(), result)
}

func (suite *LeaveTypeRepositorySuite) TestUpdate_Success() {
	expectedQuery := `UPDATE "leave_type" SET "leave_type_name"=\$1,"quota_leave"=\$2 WHERE "id" = \$3`

	leaveType := dataDummy[0]
	mockDB := suite.mocksql
	mockDB.ExpectBegin()
	mockDB.ExpectExec(expectedQuery).WithArgs(leaveType.LeaveTypeName, leaveType.QuotaLeave, leaveType.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := suite.repo.Update(leaveType)
	assert.NoError(suite.T(), err)

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func (suite *LeaveTypeRepositorySuite) TestUpdate_Failed() {
	expectedError := fmt.Errorf("failed to update leave type")

	leaveType := dataDummy[0]
	expectedQuery := `UPDATE "leave_type" SET "leave_type_name"=\$1,"quota_leave"=\$2 WHERE "id" = \$3`

	mockDB := suite.mocksql
	mockDB.ExpectBegin()
	mockDB.ExpectExec(expectedQuery).WithArgs(leaveType.LeaveTypeName, leaveType.QuotaLeave, leaveType.ID).
		WillReturnError(expectedError)
	mockDB.ExpectRollback()

	err := suite.repo.Update(leaveType)

	assert.EqualError(suite.T(), err, expectedError.Error())

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func (suite *LeaveTypeRepositorySuite) TestDelete_Success() {
	LeaveTypeID := "1"
	expectedQuery := `DELETE FROM "leave_type" WHERE id = \$1`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(LeaveTypeID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Delete(LeaveTypeID)
	assert.NoError(suite.T(), err)
}

func (suite *LeaveTypeRepositorySuite) TestDelete_Failed() {
	LeaveTypeID := "1"
	expectedError := fmt.Errorf("failed to delete leave type")

	expectedQuery := `DELETE FROM "leave_type" WHERE id = \$1`

	mockDB := suite.mocksql
	mockDB.ExpectBegin()
	mockDB.ExpectExec(expectedQuery).WithArgs(LeaveTypeID).WillReturnError(expectedError)
	mockDB.ExpectRollback()

	err := suite.repo.Delete(LeaveTypeID)

	assert.EqualError(suite.T(), err, expectedError.Error())

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func (suite *LeaveTypeRepositorySuite) TestGetByNameStatus_Success() {
	LeaveTypeName := "Annual"
	expectedQuery := `SELECT \* FROM "leave_type" WHERE leave_type_name ILIKE \$1 ORDER BY "leave_type"."id" LIMIT 1`

	rows := sqlmock.NewRows([]string{"id", "leave_type_name", "quota_leave"}).AddRow(dataDummy[0].ID, dataDummy[0].LeaveTypeName, dataDummy[0].QuotaLeave)

	suite.mocksql.ExpectQuery(expectedQuery).WithArgs("%" + LeaveTypeName + "%").WillReturnRows(rows)

	result, err := suite.repo.GetByName(LeaveTypeName)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dataDummy[0], result)

	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func (suite *LeaveTypeRepositorySuite) TestGetByNameStatus_Failed() {
	LeaveTypeName := "Annual"
	expectedError := fmt.Errorf("failed to retrieve leave type by name")

	expectedQuery := `SELECT \* FROM "leave_type" WHERE leave_type_name ILIKE \$1 ORDER BY "leave_type"."id" LIMIT 1`

	mockDB := suite.mocksql
	mockDB.ExpectQuery(expectedQuery).WithArgs("%" + LeaveTypeName + "%").WillReturnError(expectedError)

	result, err := suite.repo.GetByName(LeaveTypeName)

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.LeaveType{}, result)

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func TestLeaveTypeRepositorySuite(t *testing.T) {
	suite.Run(t, new(LeaveTypeRepositorySuite))
}
