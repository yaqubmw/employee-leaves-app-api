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

type PositionRepositorySuite struct {
	suite.Suite
	repo    PositionRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *PositionRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewPositionRepository(gormDB)
}

func (suite *PositionRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

var dataPositionDummy = []model.Position{
	{
		ID:   "1",
		Name: "Marketing",
	},
	{
		ID:   "2",
		Name: "HR",
	},
}

func (suite *PositionRepositorySuite) TestCreate_Success() {
	payload := dataPositionDummy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"position\" (.+)").WithArgs(payload.ID, payload.Name).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *PositionRepositorySuite) TestCreate_Failed() {
	payload := dataPositionDummy[0]

	expectedError := fmt.Errorf("failed to create position")

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"position\" (.+)").WithArgs(payload.ID, payload.Name).WillReturnError(expectedError)
	suite.mocksql.ExpectRollback()

	err := suite.repo.Create(payload)

	assert.EqualError(suite.T(), err, expectedError.Error())
}

func (suite *PositionRepositorySuite) TestGet_Success() {
	PositionID := "1"
	expectedPosition := dataPositionDummy[0]

	rows := sqlmock.NewRows([]string{"id", "name"})
	rows.AddRow(expectedPosition.ID, expectedPosition.Name)
	expectedQuery := `SELECT \* FROM "position" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(PositionID).WillReturnRows(rows)

	result, err := suite.repo.Get(PositionID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedPosition, result)
}

func (suite *PositionRepositorySuite) TestGet_Failed() {
	PositionID := "1"

	expectedError := fmt.Errorf("failed to retrieve position")

	expectedQuery := `SELECT \* FROM "position" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(PositionID).WillReturnError(expectedError)

	result, err := suite.repo.Get(PositionID)

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.Position{}, result)
}

func (suite *PositionRepositorySuite) TestList_Success() {
	expectedPositions := dataPositionDummy

	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, Position := range expectedPositions {
		rows.AddRow(Position.ID, Position.Name)
	}

	expectedQuery := `SELECT \* FROM "position"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

	result, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedPositions, result)
}

func (suite *PositionRepositorySuite) TestList_Failed() {
	expectedError := fmt.Errorf("failed to retrieve leave types")

	expectedQuery := `SELECT \* FROM "position"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnError(expectedError)

	result, err := suite.repo.List()

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Nil(suite.T(), result)
}

func (suite *PositionRepositorySuite) TestUpdate_Success() {
	expectedQuery := `UPDATE "position" SET "name"=\$1 WHERE "id" = \$2`

	Position := dataPositionDummy[0]
	mockDB := suite.mocksql
	mockDB.ExpectBegin()
	mockDB.ExpectExec(expectedQuery).WithArgs(Position.Name, Position.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := suite.repo.Update(Position)
	assert.NoError(suite.T(), err)

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func (suite *PositionRepositorySuite) TestUpdate_Failed() {
	expectedError := fmt.Errorf("failed to update position")

	Position := dataPositionDummy[0]
	expectedQuery := `UPDATE "position" SET "name"=\$1 WHERE "id" = \$2`

	mockDB := suite.mocksql
	mockDB.ExpectBegin()
	mockDB.ExpectExec(expectedQuery).WithArgs(Position.Name, Position.ID).
		WillReturnError(expectedError)
	mockDB.ExpectRollback()

	err := suite.repo.Update(Position)

	assert.EqualError(suite.T(), err, expectedError.Error())

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func (suite *PositionRepositorySuite) TestDelete_Success() {
	PositionID := "1"
	expectedQuery := `DELETE FROM "position" WHERE "position"."id" = \$1`

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec(expectedQuery).WithArgs(PositionID).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Delete(PositionID)
	assert.NoError(suite.T(), err)
}

func (suite *PositionRepositorySuite) TestDelete_Failed() {
	PositionID := "1"
	expectedError := fmt.Errorf("failed to delete position")

	expectedQuery := `DELETE FROM "position" WHERE "position"."id" = \$1`

	mockDB := suite.mocksql
	mockDB.ExpectBegin()
	mockDB.ExpectExec(expectedQuery).WithArgs(PositionID).WillReturnError(expectedError)
	mockDB.ExpectRollback()

	err := suite.repo.Delete(PositionID)

	assert.EqualError(suite.T(), err, expectedError.Error())

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func (suite *PositionRepositorySuite) TestGetByNameStatus_Success() {
	Name := "Marketing"
	expectedQuery := `SELECT \* FROM "position" WHERE name ILIKE \$1 ORDER BY "position"."id" LIMIT 1`

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(dataPositionDummy[0].ID, dataPositionDummy[0].Name)

	suite.mocksql.ExpectQuery(expectedQuery).WithArgs("%" + Name + "%").WillReturnRows(rows)

	result, err := suite.repo.GetByName(Name)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dataPositionDummy[0], result)

	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func (suite *PositionRepositorySuite) TestGetByNameStatus_Failed() {
	Name := "Marketing"
	expectedError := fmt.Errorf("failed to retrieve position by name")

	expectedQuery := `SELECT \* FROM "position" WHERE name ILIKE \$1 ORDER BY "position"."id" LIMIT 1`

	mockDB := suite.mocksql
	mockDB.ExpectQuery(expectedQuery).WithArgs("%" + Name + "%").WillReturnError(expectedError)

	result, err := suite.repo.GetByName(Name)

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.Position{}, result)

	assert.NoError(suite.T(), mockDB.ExpectationsWereMet())
}

func TestPositionRepositorySuite(t *testing.T) {
	suite.Run(t, new(PositionRepositorySuite))
}
