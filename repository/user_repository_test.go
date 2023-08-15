package repository

import (
	"database/sql"
	"employeeleave/model"
	"employeeleave/model/dto"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
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
		RoleId:   "",
		IsActive: true,
		Role: model.Role{
			Id:       "",
			RoleName: "",
		},
	},
	{
		ID:       "2",
		Username: "panji",
		Password: "password",
		RoleId:   "R2",
		IsActive: true,
		Role: model.Role{
			Id:       "R2",
			RoleName: "Employee",
		},
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

func (suite *UserCredentialRepositorySuite) TestGetSuccess() {
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

func (suite *UserCredentialRepositorySuite) TestUpdateSuccess() {
	expectedQuery := `UPDATE "user_credential" SET "id"=\$1,"username"=\$2,"password"=\$3,"is_active"=\$4 WHERE "id" = \$5`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(dataDummy[0].ID, dataDummy[0].Username, dataDummy[0].Password, dataDummy[0].IsActive, dataDummy[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Update(dataDummy[0])
	assert.NoError(suite.T(), err)
}

func (suite *UserCredentialRepositorySuite) TestGetByUsernameSuccess() {
	username := "agung"
	expectedQuery := `SELECT \* FROM "user_credential" WHERE username = \$1`

	rows := sqlmock.NewRows([]string{"id", "username", "password", "roleId", "is_active"}).
		AddRow(dataDummy[0].ID, dataDummy[0].Username, dataDummy[0].Password, dataDummy[0].RoleId, dataDummy[0].IsActive)

	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(username).WillReturnRows(rows)

	result, err := suite.repo.GetByUsername(username)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dataDummy[0], result)

	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func (suite *UserCredentialRepositorySuite) TestGetByUsernamePasswordSuccess() {
	// Data dummy
	username := "agung"
	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	expectedUser := model.UserCredential{
		// Isi dengan data yang diharapkan dari database
		Password: string(hashedPassword),
	}

	// Ekspektasi query GetByUsername
	suite.mocksql.ExpectQuery("^SELECT (.+) FROM \"user_credential\"*").
		WithArgs(username).
		WillReturnRows(suite.mocksql.NewRows([]string{"id", "username", "password", "roleId", "is_active"}).
			AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Password, expectedUser.RoleId, expectedUser.IsActive))

	// Panggil fungsi GetByUsernamePassword
	user, err := suite.repo.GetByUsernamePassword(username, password)

	// Assertion
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, user)

	// Verifikasi panggilan query
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func (suite *UserCredentialRepositorySuite) TestGetByUsernameFailure() {
	// Data dummy
	username := "agung"
	password := "incorrect_password"

	// Ekspektasi query GetByUsername
	suite.mocksql.ExpectQuery("^SELECT (.+) FROM \"user_credential\"*").
		WithArgs(username).
		WillReturnError(sql.ErrNoRows) // Simulasi error ketika tidak ada data yang ditemukan

	// Panggil fungsi GetByUsernamePassword
	_, err := suite.repo.GetByUsernamePassword(username, password)

	// Assertion
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// Verifikasi panggilan query
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func (suite *UserCredentialRepositorySuite) TestPagingUser_QueryPagingError() {
	suite.mocksql.ExpectQuery("^SELECT (.+) FROM \"user_credential\"*").WillReturnError(fmt.Errorf("error"))
	actualUser, actualPaging, actualError := suite.repo.Paging(dto.PaginationParam{})
	assert.Error(suite.T(), actualError)
	assert.Nil(suite.T(), actualUser)
	assert.Equal(suite.T(), actualPaging.TotalRows, 0)
}

func TestUserCredentialRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserCredentialRepositorySuite))
}
