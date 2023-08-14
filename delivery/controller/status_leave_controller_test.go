package controller

import (
	"bytes"
	"encoding/json"
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

type statusLeaveUCMock struct {
	mock.Mock
}

func (s *statusLeaveUCMock) DeleteStatusLeave(id string) error {
	args := s.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *statusLeaveUCMock) FindAllStatusLeave() ([]model.StatusLeave, error) {
	args := s.Called()
	return args.Get(0).([]model.StatusLeave), args.Error(1)
}

func (s *statusLeaveUCMock) FindByIdStatusLeave(id string) (model.StatusLeave, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return model.StatusLeave{}, args.Error(1)
	}
	return args.Get(0).(model.StatusLeave), nil
}

func (s *statusLeaveUCMock) FindByNameStatusLeave(statusName string) (model.StatusLeave, error) {
	args := s.Called(statusName)
	if args.Get(1) != nil {
		return model.StatusLeave{}, args.Error(1)
	}
	return args.Get(0).(model.StatusLeave), nil
}

func (s *statusLeaveUCMock) RegisterNewStatusLeave(payload model.StatusLeave) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *statusLeaveUCMock) UpdateStatusLeave(payload model.StatusLeave) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type StatusLeaveControllerSuite struct {
	suite.Suite
	controller  *StatusLeaveController
	useCaseMock *statusLeaveUCMock
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
}

func (suite *StatusLeaveControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.useCaseMock = &statusLeaveUCMock{}
	suite.controller = NewStatusLeaveController(suite.router, suite.useCaseMock)
	suite.recorder = httptest.NewRecorder()
}

func (suite *StatusLeaveControllerSuite) TestCreateHandler_StatusCreated() {
	payload := model.StatusLeave{
		StatusLeaveName: "Pending",
	}

	suite.useCaseMock.On("RegisterNewStatusLeave", mock.AnythingOfType("model.StatusLeave")).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/statusleaves", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
}

func (suite *StatusLeaveControllerSuite) TestCreateHandler_BadRequest() {
	payload := []model.StatusLeave{}
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/statusleaves", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *StatusLeaveControllerSuite) TestCreateHandler_InternalServerError() {
	payload := model.StatusLeave{
		StatusLeaveName: "Pending",
	}
	reqBody, _ := json.Marshal(payload)
	suite.useCaseMock.On("RegisterNewStatusLeave", mock.AnythingOfType("model.StatusLeave")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/statusleaves", bytes.NewBuffer(reqBody))
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

func (suite *StatusLeaveControllerSuite) TestListHandler_Success() {
	statusLeaves := []model.StatusLeave{
		{ID: "1", StatusLeaveName: "Pending"},
		{ID: "2", StatusLeaveName: "Approved"},
	}

	suite.useCaseMock.On("FindAllStatusLeave").Return(statusLeaves, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/statusleaves", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *StatusLeaveControllerSuite) TestListHandler_InternalServerError() {
	suite.useCaseMock.On("FindAllStatusLeave").Return([]model.StatusLeave{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/statusleaves", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *StatusLeaveControllerSuite) TestGetHandler_Success() {
	statusLeave := model.StatusLeave{
		ID:              "1",
		StatusLeaveName: "Pending",
	}

	suite.useCaseMock.On("FindByIdStatusLeave", "1").Return(statusLeave, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/statusleaves/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *StatusLeaveControllerSuite) TestGetHandler_InternalServerError() {
	suite.useCaseMock.On("FindByIdStatusLeave", "1").Return(model.StatusLeave{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/statusleaves/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

}

func (suite *StatusLeaveControllerSuite) TestUpdateHandler_StatusCreated() {
	payload := model.StatusLeave{
		ID:              "1",
		StatusLeaveName: "Pending",
	}

	reqBody, _ := json.Marshal(payload)

	suite.useCaseMock.On("UpdateStatusLeave", mock.AnythingOfType("model.StatusLeave")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/statusleaves", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *StatusLeaveControllerSuite) TestUpdateHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/statusleaves", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *StatusLeaveControllerSuite) TestUpdateHandler_InternalServerError() {
	payload := model.StatusLeave{
		StatusLeaveName: "Pending",
	}

	suite.useCaseMock.On("UpdateStatusLeave", payload).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/statusleaves", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *StatusLeaveControllerSuite) TestDeleteHandler_Success() {
	suite.useCaseMock.On("DeleteStatusLeave", "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/statusleaves/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusNoContent, suite.recorder.Code)
	assert.Empty(suite.T(), suite.recorder.Body.String())
}

func (suite *StatusLeaveControllerSuite) TestDeleteHandler_InternalServerError() {
	suite.useCaseMock.On("DeleteStatusLeave", "1").Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/statusleaves/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func TestStatusLeaveControllerSuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveControllerSuite))
}
