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

// func (suite *HistoryRepositorySuite) TestList(requestPaging dto.PaginationParam) ([]model.HistoryLeave, dto.Paging, error) {
// 	expectedHistories := dataDummyHistory
// 	expectedPaging := dto.Paging{
// 		Page:        1,
// 		RowsPerPage: 5,
// 		TotalRows:   5,
// 		TotalPages:  1,
// 	}
// 	requestPaging := dto.PaginationParam{
// 		Page:   1,
// 		Limit:  5,
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "transaction_leave_id", "date_event"})
// 	for _, history := range expectedHistories {
// 		rows.AddRow(history.Id, history.TransactionLeaveId, history.DateEvent)
// 	}

// 	expectedQuery := `SELECT \* FROM "history_leave"`
// 	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

// 	result, err := suite.repo.Paging("Paging", requestPaging)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedHistories, result)
// }

// func (suite *HistoryRepositorySuite) TestUpdate() {
// 	payload := dataDummy[0]

// 	expectedQuery := `UPDATE "role" SET "id"=$1,"role_name"=$2 WHERE "id" = $3`


// 	suite.mocksql.ExpectExec(expectedQuery).WithArgs(payload.TransactionLeaveId, payload.Id).WillReturnResult(sqlmock.NewResult(1, 1))


// 	err := suite.repo.Update(payload)
// 	assert.NoError(suite.T(), err)

// 	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
// }

func TestHistoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(HistoryRepositorySuite))
}
