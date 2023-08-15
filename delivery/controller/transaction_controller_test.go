package controller

import (
	"bytes"
	"employeeleave/model"
	"employeeleave/model/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type transactionUCMock struct {
	mock.Mock
}

func (t *transactionUCMock) FindAllEmpl(params dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error) {
	args := t.Called(params)
	return args.Get(0).([]dto.TransactionResponseDto), args.Get(1).(dto.Paging), args.Error(2)
}

func (t *transactionUCMock) FindById(id string) (model.TransactionLeave, error) {
	args := t.Called(id)
	if args.Get(1) != nil {
		return model.TransactionLeave{}, args.Error(1)
	}
	return args.Get(0).(model.TransactionLeave), nil
}

func (t *transactionUCMock) FindByIdEmpl(id string) ([]model.TransactionLeave, error) {
	args := t.Called(id)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TransactionLeave), nil
}

func (t *transactionUCMock) ApplyLeave(payload model.TransactionLeave) error {
	args := t.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (t *transactionUCMock) ApproveOrRejectLeave(payload model.TransactionLeave) error {
	args := t.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type TransactionControllerSuite struct {
	suite.Suite
	controller  *TransactionLeaveController
	useCaseMock *transactionUCMock
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
}

func (suite *TransactionControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.useCaseMock = &transactionUCMock{}
	suite.controller = NewTransactionController(suite.router, suite.useCaseMock)
	suite.recorder = httptest.NewRecorder()
}

func (suite *TransactionControllerSuite) TestCreateHandler_StatusCreated() {
	payload := model.TransactionLeave{
		ID:             "1",
		EmployeeID:     "employee1",
		LeaveTypeID:    "leaveType1",
		StatusLeaveID:  "status1",
		DateStart:      time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC),
		DateEnd:        time.Date(2023, time.August, 16, 0, 0, 0, 0, time.UTC),
		Reason:         "Vacation",
		SubmissionDate: time.Date(2023, time.August, 14, 10, 0, 0, 0, time.UTC),
		AmountLeave:    1,
		TypeOfDay:      "Full Day",
	}

	// Set up the mock to expect ApplyLeave method call and return nil (no error)
	suite.useCaseMock.On("ApplyLeave", payload).Return(nil)

	// Prepare the request with the payload
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is StatusCreated
	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
}

func (suite *TransactionControllerSuite) TestCreateHandler_BadRequest() {
	// Create an invalid payload (empty array)
	payload := []model.TransactionLeave{}

	// Prepare the request with the payload
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is BadRequest
	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *TransactionControllerSuite) TestCreateHandler_InternalServerError() {
	// Create a valid payload
	payload := model.TransactionLeave{
		ID: "1",
	}

	// Prepare the request with the payload
	reqBody, _ := json.Marshal(payload)

	// Mock the ApplyLeave function to return an error
	suite.useCaseMock.On("ApplyLeave", mock.AnythingOfType("model.TransactionLeave")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is InternalServerError
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

	// Assert the response body contains the expected error message
	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string
	}
	json.Unmarshal(response, &actualError)
	assert.Equal(suite.T(), "error", actualError.Err)
}

func (suite *TransactionControllerSuite) TestListHandler_Success() {
	// Create mock data for transactions
	transactions := []dto.TransactionResponseDto{
		{ID: "1"},
		{ID: "2"},
	}

	// Mock the FindAllEmpl function to return the mock data
	suite.useCaseMock.On("FindAllEmpl", mock.AnythingOfType("dto.PaginationParam")).Return(transactions, dto.Paging{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/transaction", nil)

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is OK
	require.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

	// Parse the response body into an array of transaction DTOs
	var responseBody []dto.TransactionResponseDto
	err := json.Unmarshal(suite.recorder.Body.Bytes(), &responseBody)
	require.Nil(suite.T(), err)

	// Assert that the number of response DTOs matches the number of mock transactions
	require.Equal(suite.T(), len(transactions), len(responseBody))

	// Assert that the IDs of the response DTOs match the mock data IDs
	for i, mockTransaction := range transactions {
		require.Equal(suite.T(), mockTransaction.ID, responseBody[i].ID)
	}
}

func (suite *TransactionControllerSuite) TestListHandler_InternalServerError() {
	// Mock the FindAllEmpl function to return an error
	suite.useCaseMock.On("FindAllEmpl", mock.AnythingOfType("dto.PaginationParam")).Return(nil, dto.Paging{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/transaction", nil)

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is Internal Server Error
	require.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

	// Parse the response body to check the error message
	var responseBody struct {
		Err string `json:"err"`
	}
	err := json.Unmarshal(suite.recorder.Body.Bytes(), &responseBody)
	require.Nil(suite.T(), err)

	// Assert that the error message matches the expected error message
	require.Equal(suite.T(), "error", responseBody.Err)
}

func (suite *TransactionControllerSuite) TestGetHandler_Success() {
	// Create a mock transaction object
	mockTransaction := model.TransactionLeave{
		ID: "1",
		// ... other fields ...
	}

	// Mock the FindByIdEmpl function to return the mockTransaction
	suite.useCaseMock.On("FindByIdEmpl", "1").Return(mockTransaction, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/transaction/1", nil)

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is OK
	require.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

	// Parse the response body to check the transaction details
	var responseBody model.TransactionLeave
	err := json.Unmarshal(suite.recorder.Body.Bytes(), &responseBody)
	require.Nil(suite.T(), err)

	// Assert that the response transaction matches the mockTransaction
	require.Equal(suite.T(), mockTransaction, responseBody)
}

func (suite *TransactionControllerSuite) TestGetHandler_InternalServerError() {
	// Mock the FindByIdEmpl function to return an error
	suite.useCaseMock.On("FindByIdEmpl", "1").Return(model.TransactionLeave{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/transaction/1", nil)

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is Internal Server Error
	require.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

	// Parse the response body to check the error details
	var responseBody struct {
		Err string `json:"err"`
	}
	err := json.Unmarshal(suite.recorder.Body.Bytes(), &responseBody)
	require.Nil(suite.T(), err)

	// Assert that the error message matches the expected error message
	require.Equal(suite.T(), "error", responseBody.Err)
}

func (suite *TransactionControllerSuite) TestUpdateHandler_StatusOK() {
	payload := model.TransactionLeave{
		ID: "1",
	}

	// Mock the UpdateEmpl function to return no error
	suite.useCaseMock.On("UpdateEmpl", mock.AnythingOfType("model.TransactionLeave")).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is OK
	require.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *TransactionControllerSuite) TestUpdateHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/transaction", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is BadRequest
	require.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *TransactionControllerSuite) TestUpdateHandler_InternalServerError() {
	payload := model.TransactionLeave{
		ID: "1",
	}

	suite.useCaseMock.On("UpdateEmpl", mock.AnythingOfType("model.Transaction_Leave")).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Mock the use case to return an error
	suite.useCaseMock.On("UpdateEmpl", mock.AnythingOfType("model.Transaction_Leave")).Return(fmt.Errorf("error"))

	// Serve the request and capture the response
	suite.router.ServeHTTP(suite.recorder, req)

	// Assert the response status code is InternalServerError
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func TestTransactionControllerSuite(t *testing.T) {
	suite.Run(t, new(TransactionControllerSuite))
}
