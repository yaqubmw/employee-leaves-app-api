package repository

import (
	"employeeleave/model"
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
// func (suite *HistoryRepositorySuite) TestPaging() {
// 	// err := godotenv.Load("../.env") // Adjust the path to your .env file
// 	// if err != nil {
// 	// 	suite.T().Fatal("Error loading .env file:", err)
// 	// }

// 	page := 1
// 	perPage := 10
// 	expectedHistories := dataDummyHistory

// 	// Mock data for pagination
// 	rows := sqlmock.NewRows([]string{"id", "transaction_leave_id", "date_event"})
// 	for _, history := range expectedHistories {
// 		rows.AddRow(history.Id, history.TransactionLeaveId, history.DateEvent)
// 	}
// 	expectedQuery := `SELECT \* FROM "history_leave" LIMIT \$1 OFFSET \$2`
// 	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(perPage, (page-1)*perPage).WillReturnRows(rows)

// 	// Create a PaginationParam
// 	paginationParam := dto.PaginationParam{
// 		Page:   page,
// 		Offset: (page - 1) * perPage,
// 		Limit:  perPage,
// 	}

// 	// Call the Paging method and assert the results
// 	results, _, err := suite.repo.Paging(paginationParam)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedHistories, results)
// }

func TestHistoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(HistoryRepositorySuite))
}
