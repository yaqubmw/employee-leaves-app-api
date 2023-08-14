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

type StatusLeaveRepositorySuite struct {
	suite.Suite
	repo    StatusLeaveRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *StatusLeaveRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewStatusLeaveRepository(gormDB)
}

func (suite *StatusLeaveRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

var statusDummy = []model.StatusLeave{
	{
		ID:              "1",
		StatusLeaveName: "Pending",
	},
	{
		ID:              "2",
		StatusLeaveName: "Approved",
	},
}

func (suite *StatusLeaveRepositorySuite) TestCreate() {
	payload := statusDummy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"status_leave\" (.+)").WithArgs(payload.ID, payload.StatusLeaveName).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *StatusLeaveRepositorySuite) TestGet() {
	statusLeaveID := "1"
	expectedStatusLeave := statusDummy[0]

	rows := sqlmock.NewRows([]string{"id", "status_leave_name"})
	rows.AddRow(expectedStatusLeave.ID, expectedStatusLeave.StatusLeaveName)
	expectedQuery := `SELECT \* FROM "status_leave" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(statusLeaveID).WillReturnRows(rows)

	result, err := suite.repo.Get(statusLeaveID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatusLeave, result)
}

func (suite *StatusLeaveRepositorySuite) TestList() {
	expectedStatusLeaves := statusDummy

	rows := sqlmock.NewRows([]string{"id", "status_leave_name"})
	for _, statusLeave := range expectedStatusLeaves {
		rows.AddRow(statusLeave.ID, statusLeave.StatusLeaveName)
	}

	expectedQuery := `SELECT \* FROM "status_leave"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

	result, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatusLeaves, result)
}

func (suite *StatusLeaveRepositorySuite) TestUpdate() {
	expectedQuery := `UPDATE "status_leave" SET "id"=\$1,"status_leave_name"=\$2 WHERE "id" = \$3`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(statusDummy[0].ID, statusDummy[0].StatusLeaveName, statusDummy[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Update(statusDummy[0])
	assert.NoError(suite.T(), err)
}

func (suite *StatusLeaveRepositorySuite) TestDelete() {
	statusLeaveID := "1"
	expectedQuery := `DELETE FROM "status_leave" WHERE id = \$1`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(statusLeaveID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Delete(statusLeaveID)
	assert.NoError(suite.T(), err)
}

func (suite *StatusLeaveRepositorySuite) TestGetByNameStatus() {
	statusLeaveName := "Pending"
	expectedQuery := `SELECT \* FROM "status_leave" WHERE status_leave_name LIKE \$1`

	rows := sqlmock.NewRows([]string{"id", "status_leave_name"}).AddRow(statusDummy[0].ID, statusDummy[0].StatusLeaveName)

	suite.mocksql.ExpectQuery(expectedQuery).WithArgs("%" + statusLeaveName + "%").WillReturnRows(rows)

	result, err := suite.repo.GetByNameStatus(statusLeaveName)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), statusDummy[0], result)

	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func TestStatusLeaveRepositorySuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveRepositorySuite))
}
