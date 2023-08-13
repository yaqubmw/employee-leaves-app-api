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

type RoleRepositorySuite struct {
	suite.Suite
	repo    RoleRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *RoleRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewRoleRepository(gormDB)
}

func (suite *RoleRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

var dataDummy = []model.Role{
	{
		Id:              "1",
		RoleName: "Pending",
	},
	{
		Id:              "2",
		RoleName: "Approved",
	},
}

func (suite *RoleRepositorySuite) TestCreate() {
	payload := dataDummy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"role\" (.+)").WithArgs(payload.Id, payload.RoleName).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *RoleRepositorySuite) TestGet() {
	roleId := "1"
	expectedRole := dataDummy[0]

	rows := sqlmock.NewRows([]string{"id", "role_name"})
	rows.AddRow(expectedRole.Id, expectedRole.RoleName)
	expectedQuery := `SELECT \* FROM "role" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(roleId).WillReturnRows(rows)

	result, err := suite.repo.Get(roleId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRole, result)
}

func (suite *RoleRepositorySuite) TestList() {
	expectedRoles := dataDummy

	rows := sqlmock.NewRows([]string{"id", "role_name"})
	for _, role := range expectedRoles {
		rows.AddRow(role.Id, role.RoleName)
	}

	expectedQuery := `SELECT \* FROM "role"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

	result, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoles, result)
}

// func (suite *RoleRepositorySuite) TestUpdate() {
// 	payload := dataDummy[0]

// 	expectedQuery := `UPDATE "role" SET "id"=$1,"role_name"=$2 WHERE "id" = $3`


// 	suite.mocksql.ExpectExec(expectedQuery).WithArgs(payload.RoleName, payload.Id).WillReturnResult(sqlmock.NewResult(1, 1))


// 	err := suite.repo.Update(payload)
// 	assert.NoError(suite.T(), err)

// 	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
// }

func TestRoleRepositorySuite(t *testing.T) {
	suite.Run(t, new(RoleRepositorySuite))
}
