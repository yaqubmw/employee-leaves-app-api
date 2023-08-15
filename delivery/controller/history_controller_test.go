package controller

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type historyUseCaseMock struct {
	mock.Mock
}

// FindAllHistory implements usecase.HistoryUseCase.
func (h *historyUseCaseMock) FindAllHistory(requestPaging dto.PaginationParam) ([]model.HistoryLeave, dto.Paging, error) {
	args := h.Called(requestPaging)
	return args.Get(0).([]model.HistoryLeave), args.Get(1).(dto.Paging), args.Error(2)
}

// FindHistoryById implements usecase.HistoryUseCase.
func (h *historyUseCaseMock) FindHistoryById(id string) (model.HistoryLeave, error) {
	args := h.Called(id)
	return args.Get(0).(model.HistoryLeave), args.Error(1)
}

// RegisterNewHistory implements usecase.HistoryUseCase.
func (h *historyUseCaseMock) RegisterNewHistory(payload model.HistoryLeave) error {
	args := h.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type HistoryControllerTestSuite struct {
	suite.Suite
	usecaseMock *historyUseCaseMock
	router      *gin.Engine
	controller  *HistoryController
	recorder    *httptest.ResponseRecorder
}

func (suite *HistoryControllerTestSuite) SetupTest() {
	suite.usecaseMock = new(historyUseCaseMock)
	suite.router = gin.Default()
	suite.recorder = httptest.NewRecorder()
	suite.controller = NewHistoryController(suite.router, suite.usecaseMock)
}

func TestHistoryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(HistoryControllerTestSuite))
}

func parseTime(timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return parsedTime
}

func (suite *HistoryControllerTestSuite) TestGetHandler_Success() {
	recorder := httptest.NewRecorder()
	history := model.HistoryLeave{
		Id:          "1",
		TransactionLeaveId: "1",
		DateEvent: parseTime("2023-08-12T10:30:00Z"),
	}

	suite.usecaseMock.On("FindHistoryById", "1").Return(history, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/histories/1", nil)
	suite.router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *HistoryControllerTestSuite) TestGetHandler_InternalServerError() {
	suite.usecaseMock.On("FindHistoryById", "1").Return(model.HistoryLeave{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/histories/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

}

func (suite *HistoryControllerTestSuite) TestListHandler_Success() {
	expectedPaginationParam := dto.PaginationParam{
		Page: 1,
		Limit: 5,
	}

	expectedHistory := []model.HistoryLeave{{
		Id:          "1",
		TransactionLeaveId: "1",
		DateEvent: parseTime("2023-08-12T10:30:00Z"),
	}}

	expectedPaging := dto.Paging{
		Page: 1,
		RowsPerPage: 5,
		TotalRows: 1,
		TotalPages: 1,
	}

	suite.usecaseMock.On("FindAllHistory", expectedPaginationParam).Return(expectedHistory, expectedPaging, nil)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/histories?page=1&limit=5", nil)
	suite.router.ServeHTTP(recorder, request)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}


func (suite *HistoryControllerTestSuite) TestListHandler_InternalServerError() {
	// Mock the usecase's FindAllHistory function to return an error
	suite.usecaseMock.On("FindAllHistory", dto.PaginationParam{Page: 1, Limit: 5}).Return(nil, dto.Paging{}, errors.New("error"))

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/histories?page=1&limit=5", nil)
	suite.router.ServeHTTP(suite.recorder, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

	// Parse the response body only if it's not empty
	if len(suite.recorder.Body.Bytes()) > 0 {
		var actualError struct {
			Err string `json:"error"`
		}
		err := json.Unmarshal(suite.recorder.Body.Bytes(), &actualError)
		require.NoError(suite.T(), err) // Check that unmarshaling was successful

		assert.Equal(suite.T(), "error", actualError.Err)
	}
}
