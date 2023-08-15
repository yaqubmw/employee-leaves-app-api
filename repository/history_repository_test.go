package repository

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HistoryRepositorySuite struct {
	suite.Suite
	repo    HistoryRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *HistoryRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewHistoryRepository(gormDB)
}

func (suite *HistoryRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func parseTime(timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return parsedTime
}

var dataDummyHistory = []model.HistoryLeave{
	{
		Id:              "1",
		TransactionLeaveId: "Pending",
		DateEvent: parseTime("2023-08-12T10:30:00Z"),
	},
	{
		Id:              "2",
		TransactionLeaveId: "Approved",
		DateEvent: parseTime("2023-08-12T10:30:00Z"),
	},
}

func (suite *HistoryRepositorySuite) TestCreate() {
	payload := dataDummyHistory[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"history_leave\" (.+)").WithArgs(payload.Id, payload.TransactionLeaveId, payload.DateEvent).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *HistoryRepositorySuite) TestGet() {
	historyId := "1"
	expectedHistory := dataDummyHistory[0]

	rows := sqlmock.NewRows([]string{"id", "transaction_leave_id", "date_event"})
	rows.AddRow(expectedHistory.Id, expectedHistory.TransactionLeaveId, expectedHistory.DateEvent)
	expectedQuery := `SELECT \* FROM "history_leave" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(historyId).WillReturnRows(rows)

	result, err := suite.repo.GetHistoryById(historyId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedHistory, result)
}

// TEST FAIL : Error loading .env file
func (suite *HistoryRepositorySuite) TestPagingHistory_QueryPagingError() {
	suite.mocksql.ExpectQuery("^SELECT (.+) FROM \"history_leave\"*").WillReturnError(fmt.Errorf("error"))
	actualHistory, actualPaging, actualError := suite.repo.Paging(dto.PaginationParam{})
	assert.Error(suite.T(), actualError)
	assert.Nil(suite.T(), actualHistory)
	assert.Equal(suite.T(), actualPaging.TotalRows, 0)
}

func TestHistoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(HistoryRepositorySuite))
}
