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

type UserCredentialRepositorySuite struct {
	suite.Suite
	repo    UserRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *UserCredentialRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewUserRepository(gormDB)
}

func (suite *UserCredentialRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

var dataDummy = []model.UserCredential{
	{
		ID:       "1",
		Username: "agung",
		Password: "password",
		RoleId:   "R1",
		IsActive: true,
	},
	{
		ID:       "1",
		Username: "panji",
		Password: "password",
		RoleId:   "R2",
		IsActive: true,
	},
}

func (suite *UserCredentialRepositorySuite) TestCreate() {
	payload := dataDummy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"user_credential\" (.+)").WithArgs(payload.ID, payload.Username, payload.Password, payload.RoleId, payload.IsActive).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *UserCredentialRepositorySuite) TestGet() {
	userID := "1"
	expectedUser := dataDummy[0]

	rows := sqlmock.NewRows([]string{"id", "username", "password", "roleId", "is_active"})
	rows.AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Password, expectedUser.RoleId, expectedUser.IsActive)
	expectedQuery := `SELECT \* FROM "user_credential" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(userID).WillReturnRows(rows)

	result, err := suite.repo.Get(userID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, result)
}

// func (suite *UserCredentialRepositorySuite) TestList() {
// 	expectedStatusLeaves := dataDummy

// 	rows := sqlmock.NewRows([]string{"id", "status_leave_name"})
// 	for _, statusLeave := range expectedStatusLeaves {
// 		rows.AddRow(statusLeave.ID, statusLeave.StatusLeaveName)
// 	}

// 	expectedQuery := `SELECT \* FROM "status_leave"`
// 	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

// 	result, err := suite.repo.List()
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedStatusLeaves, result)
// }

// func (suite *UserCredentialRepositorySuite) TestUpdate() {
// 	expectedQuery := `UPDATE "status_leave" SET "id"=\$1,"status_leave_name"=\$2 WHERE "id" = \$3`

// 	suite.mocksql.ExpectBegin()
// 	suite.mocksql.ExpectExec(expectedQuery).WithArgs(dataDummy[0].ID, dataDummy[0].StatusLeaveName, dataDummy[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
// 	suite.mocksql.ExpectCommit()

// 	err := suite.repo.Update(dataDummy[0])
// 	assert.NoError(suite.T(), err)
// }

// func (suite *UserCredentialRepositorySuite) TestDelete() {
// 	statusLeaveID := "1"
// 	expectedQuery := `DELETE FROM "status_leave" WHERE id = \$1`

// 	suite.mocksql.ExpectBegin()
// 	suite.mocksql.ExpectExec(expectedQuery).WithArgs(statusLeaveID).WillReturnResult(sqlmock.NewResult(0, 1))
// 	suite.mocksql.ExpectCommit()

// 	err := suite.repo.Delete(statusLeaveID)
// 	assert.NoError(suite.T(), err)
// }

// func (suite *UserCredentialRepositorySuite) TestGetByNameStatus() {
// 	statusLeaveName := "Pending"
// 	expectedQuery := `SELECT \* FROM "status_leave" WHERE status_leave_name LIKE \$1`

// 	rows := sqlmock.NewRows([]string{"id", "status_leave_name"}).AddRow(dataDummy[0].ID, dataDummy[0].StatusLeaveName)

// 	suite.mocksql.ExpectQuery(expectedQuery).WithArgs("%" + statusLeaveName + "%").WillReturnRows(rows)

// 	result, err := suite.repo.GetByNameStatus(statusLeaveName)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), dataDummy[0], result)

// 	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
// }

func TestUserCredentialRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserCredentialRepositorySuite))
}
