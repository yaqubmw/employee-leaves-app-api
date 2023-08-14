package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"employeeleave/model"
)

type leaveTypeUCMock struct {
	mock.Mock
}

func (s *leaveTypeUCMock) DeleteLeaveType(id string) error {
	args := s.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *leaveTypeUCMock) FindAllLeaveType() ([]model.LeaveType, error) {
	args := s.Called()
	return args.Get(0).([]model.LeaveType), args.Error(1)
}

func (s *leaveTypeUCMock) FindByIdLeaveType(id string) (model.LeaveType, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return model.LeaveType{}, args.Error(1)
	}
	return args.Get(0).(model.LeaveType), nil
}

func (s *leaveTypeUCMock) GetByName(leave_type_name string) (model.LeaveType, error) {
	args := s.Called(leave_type_name)
	if args.Get(1) != nil {
		return model.LeaveType{}, args.Error(1)
	}
	return args.Get(0).(model.LeaveType), nil
}

func (s *leaveTypeUCMock) RegisterNewLeaveType(payload model.LeaveType) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *leaveTypeUCMock) UpdateLeaveType(payload model.LeaveType) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type LeaveTypeControllerSuite struct {
	suite.Suite
	controller  *LeaveTypeController
	useCaseMock *leaveTypeUCMock
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
}

func (suite *LeaveTypeControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.useCaseMock = &leaveTypeUCMock{}
	suite.controller = NewLeaveTypeController(suite.useCaseMock, suite.router)
	suite.recorder = httptest.NewRecorder()
}

func (suite *LeaveTypeControllerSuite) TestCreateHandler_LeaveTypeCreated() {
	payload := model.LeaveType{
		LeaveTypeName: "Annual Leave",
	}

	suite.useCaseMock.On("RegisterNewLeaveType", mock.AnythingOfType("model.LeaveType")).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/leavetypes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestCreateHandler_BadRequest() {
	payload := []model.LeaveType{}
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/leavetypes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestCreateHandler_InternalServerError() {
	payload := model.LeaveType{
		LeaveTypeName: "Annual Leave",
	}
	reqBody, _ := json.Marshal(payload)
	suite.useCaseMock.On("RegisterNewLeaveType", mock.AnythingOfType("model.LeaveType")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/leavetypes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)
	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestListLeaveTypeHandler_Success() {
	leaveTypes := []model.LeaveType{
		{ID: "1", LeaveTypeName: "Annual Leave"},
		{ID: "2", LeaveTypeName: "Maternity Leave"},
	}

	suite.useCaseMock.On("FindAllLeaveType").Return(leaveTypes, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leavetypes", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *LeaveTypeControllerSuite) TestListLeaveTypeHandler_InternalServerError() {
	suite.useCaseMock.On("FindAllLeaveType").Return([]model.LeaveType{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leavetypes", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestGetLeaveTypeHandler_Success() {
	leaveType := model.LeaveType{
		ID:            "1",
		LeaveTypeName: "Annual Leave",
		QuotaLeave:    12,
	}

	suite.useCaseMock.On("FindByIdLeaveType", "1").Return(leaveType, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leavetypes/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *LeaveTypeControllerSuite) TestGetLeaveTypeHandler_InternalServerError() {
	suite.useCaseMock.On("FindByIdLeaveType", "1").Return(model.LeaveType{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leavetypes/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

}

func (suite *LeaveTypeControllerSuite) TestGetByNameLeaveTypeHandler_Success() {
	expectedLeaveType := model.LeaveType{
		LeaveTypeName: "Annual Leave",
	}

	suite.useCaseMock.On("GetByName", "Annual Leave").Return(expectedLeaveType, nil)

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/leavetypes/name/Annual Leave", nil)
	suite.router.ServeHTTP(suite.recorder, request)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

	var response struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Data model.LeaveType `json:"data"`
	}

	err := json.Unmarshal(suite.recorder.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), 200, response.Status.Code)
	assert.Equal(suite.T(), "Get By Name Data Successfully", response.Status.Description)
	assert.Equal(suite.T(), expectedLeaveType, response.Data)
}

func (suite *LeaveTypeControllerSuite) TestGetByNameLeaveTypeHandler_InternalServerError() {
	suite.useCaseMock.On("GetByName", "Annual Leave").Return(model.LeaveType{}, errors.New("internal server error"))

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/leavetypes/name/Annual Leave", nil)
	suite.router.ServeHTTP(suite.recorder, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

	var response struct {
		Err string `json:"err"`
	}

	err := json.Unmarshal(suite.recorder.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "internal server error", response.Err)
}

func (suite *LeaveTypeControllerSuite) TestUpdateLeaveTypeHandler_LeaveTypeCreated() {
	payload := model.LeaveType{
		ID:            "1",
		LeaveTypeName: "Annual Leave",
		QuotaLeave:    12,
	}

	reqBody, _ := json.Marshal(payload)

	suite.useCaseMock.On("UpdateLeaveType", mock.AnythingOfType("model.LeaveType")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/leavetypes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestUpdateLeaveTypeHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/leavetypes", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestUpdateLeaveTypeHandler_InternalServerError() {
	payload := model.LeaveType{
		LeaveTypeName: "Annual Leave",
	}

	suite.useCaseMock.On("UpdateLeaveType", payload).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/leavetypes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *LeaveTypeControllerSuite) TestDeleteHandler_Success() {
	suite.useCaseMock.On("DeleteLeaveType", "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/leavetypes/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusNoContent, suite.recorder.Code)
	assert.Empty(suite.T(), suite.recorder.Body.String())
}

func (suite *LeaveTypeControllerSuite) TestDeleteHandler_InternalServerError() {
	suite.useCaseMock.On("DeleteLeaveType", "1").Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/leavetypes/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func TestLeaveTypeControllerSuite(t *testing.T) {
	suite.Run(t, new(LeaveTypeControllerSuite))
}
