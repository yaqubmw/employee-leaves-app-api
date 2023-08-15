package repository

// import (
// 	"employeeleave/model"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type TransactionRepositorySuite struct {
// 	suite.Suite
// 	repo    transactionRepository
// 	mockDB  *gorm.DB
// 	mocksql sqlmock.Sqlmock
// }

// func (suite *TransactionRepositorySuite) SetupTest() {
// 	db, mock, _ := sqlmock.New()
// 	gormDB, _ := gorm.Open(
// 		postgres.New(postgres.Config{
// 			Conn: db,
// 		}),
// 		&gorm.Config{},
// 	)
// 	suite.mockDB = gormDB
// 	suite.mocksql = mock
// 	suite.repo = transactionRepository{db: gormDB}
// }

// func (suite *TransactionRepositorySuite) TearDownTest() {
// 	if suite.mocksql != nil {
// 		assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
// 	}
// }

// var TransactionDumy = []model.TransactionLeave{
// 	{
// 		ID:             "1",
// 		EmployeeID:     "1",
// 		LeaveTypeID:    "1",
// 		StatusLeaveID:  "1",
// 		DateStart:      time.Now(),
// 		DateEnd:        time.Now().Add(time.Hour * 24),
// 		Reason:         "Vacation",
// 		SubmissionDate: time.Now(),
// 		AmountLeave:    1,
// 		TypeOfDay:      "Full Day",
// 	},
// 	{
// 		ID:             "2",
// 		EmployeeID:     "2",
// 		LeaveTypeID:    "2",
// 		StatusLeaveID:  "2",
// 		DateStart:      time.Now(),
// 		DateEnd:        time.Now().Add(time.Hour * 24),
// 		Reason:         "Sick",
// 		SubmissionDate: time.Now(),
// 		AmountLeave:    2,
// 		TypeOfDay:      "Full Day",
// 	},
// }

// func (suite *TransactionRepositorySuite) TestCreateTransactionBeginSuccess() {
// 	payload := TransactionDumy[0]

// 	// Expect a transaction begin
// 	suite.mocksql.ExpectBegin()

// 	// Expect the Create query
// 	suite.mocksql.ExpectExec(`INSERT INTO "transaction_leave" \(.+\)`).
// 		WithArgs(
// 			payload.ID,
// 			payload.EmployeeID,
// 			payload.LeaveTypeID,
// 			payload.StatusLeaveID,
// 			payload.DateStart,
// 			payload.DateEnd,
// 			payload.Reason,
// 			payload.SubmissionDate,
// 			payload.AmountLeave,
// 			payload.TypeOfDay,
// 		).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// Expect a transaction commit
// 	suite.mocksql.ExpectCommit()

// 	err := suite.repo.Create(payload)
// 	assert.NoError(suite.T(), err)
// }

// func (suite *TransactionRepositorySuite) TestCreateTransactionBeginError() {
// 	payload := TransactionDumy[0]

// 	// Simulate an error during the transaction begin
// 	suite.mocksql.ExpectBegin().WillReturnError(fmt.Errorf("begin error"))

// 	err := suite.repo.Create(payload)
// 	assert.Error(suite.T(), err)
// 	assert.Contains(suite.T(), err.Error(), "begin error")
// }

// func (suite *TransactionRepositorySuite) TestGetByIDSuccess() {
// 	transactionID := "1"
// 	expectedTransaction := model.TransactionLeave{
// 		ID: "1",
// 		// ... set other fields accordingly
// 	}
// 	rows := sqlmock.NewRows([]string{"id", "employee_id", "leave_type_id", "status_leave_id"}).
// 		AddRow(expectedTransaction.ID, expectedTransaction.EmployeeID, expectedTransaction.LeaveTypeID, expectedTransaction.StatusLeaveID)

// 	suite.mocksql.ExpectQuery(`SELECT \* FROM "transaction_leave" WHERE id = \$1`).
// 		WithArgs(transactionID).
// 		WillReturnRows(rows)

// 	result, err := suite.repo.GetByID(transactionID)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedTransaction, result)
// }

// func (suite *TransactionRepositorySuite) TestGetByEmployeeIDSuccess() {
// 	employeeID := "1"
// 	expectedTransaction := model.TransactionLeave{
// 		ID: "1",
// 		// ... set other fields accordingly
// 	}
// 	rows := sqlmock.NewRows([]string{"id", "employee_id", "leave_type_id", "status_leave_id"}).
// 		AddRow(expectedTransaction.ID, expectedTransaction.EmployeeID, expectedTransaction.LeaveTypeID, expectedTransaction.StatusLeaveID)

// 	suite.mocksql.ExpectQuery(`SELECT \* FROM "transaction_leave" WHERE employee_id = \$1`).
// 		WithArgs(employeeID).
// 		WillReturnRows(rows)

// 	results, err := suite.repo.GetByEmployeeID(employeeID)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), []model.TransactionLeave{expectedTransaction}, results)
// }

// func TestTransactionRepositorySuite(t *testing.T) {
// 	suite.Run(t, new(TransactionRepositorySuite))
// }
