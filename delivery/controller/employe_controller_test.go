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

type employeeUCMock struct {
	mock.Mock
}

// PaternityLeave implements usecase.EmployeeUseCase.
func (*employeeUCMock) PaternityLeave(id string, availableDays int) error {
	panic("unimplemented")
}

// UpdateAnnualLeave implements usecase.EmployeeUseCase.
func (*employeeUCMock) UpdateAnnualLeave(id string, availableDays int) error {
	panic("unimplemented")
}

// UpdateMarriageLeave implements usecase.EmployeeUseCase.
func (*employeeUCMock) UpdateMarriageLeave(id string, availableDays int) error {
	panic("unimplemented")
}

// UpdateMaternityLeave implements usecase.EmployeeUseCase.
func (*employeeUCMock) UpdateMaternityLeave(id string, availableDays int) error {
	panic("unimplemented")
}

// UpdateMenstrualLeave implements usecase.EmployeeUseCase.
func (*employeeUCMock) UpdateMenstrualLeave(id string, availableDays int) error {
	panic("unimplemented")
}

func (e *employeeUCMock) DeleteEmpl(id string) error {
	args := e.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (e *employeeUCMock) FindAllEmpl() ([]model.Employee, error) {
	args := e.Called()
	return args.Get(0).([]model.Employee), args.Error(1)
}

func (e *employeeUCMock) FindByIdEmpl(id string) (model.Employee, error) {
	args := e.Called(id)
	if args.Get(1) != nil {
		return model.Employee{}, args.Error(1)
	}
	return args.Get(0).(model.Employee), nil
}

func (e *employeeUCMock) FindByNameEmpl(name string) (model.Employee, error) {
	args := e.Called(name)
	if args.Get(1) != nil {
		return model.Employee{}, args.Error(1)
	}
	return args.Get(0).(model.Employee), nil
}

func (e *employeeUCMock) RegisterNewEmpl(payload model.Employee) error {
	args := e.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (e *employeeUCMock) UpdateEmpl(payload model.Employee) error {
	args := e.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type EmployeeControllerSuite struct {
	suite.Suite
	controller  *EmployeeController
	useCaseMock *employeeUCMock
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
}

func (suite *EmployeeControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.useCaseMock = &employeeUCMock{}
	suite.controller = NewEmplController(suite.router, suite.useCaseMock)
	suite.recorder = httptest.NewRecorder()
}

func (suite *EmployeeControllerSuite) TestCreateHandler_StatusCreated() {
	payload := model.Employee{
		ID:   "1",
		Name: "imron",
	}

	suite.useCaseMock.On("RegisterNewEmpl", mock.AnythingOfType("model.Employee")).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
}

func (suite *EmployeeControllerSuite) TestCreateHandler_BadRequest() {
	payload := []model.Employee{}
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *EmployeeControllerSuite) TestCreateHandler_InternalServerError() {
	payload := model.Employee{
		ID:   "1",
		Name: "imron",
	}
	reqBody, _ := json.Marshal(payload)
	suite.useCaseMock.On("RegisterNewEmpl", mock.AnythingOfType("model.Employee")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee", bytes.NewBuffer(reqBody))
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

func (suite *EmployeeControllerSuite) TestListHandler_Success() {
	employee := []model.Employee{
		{ID: "1", Name: "imron"},
		{ID: "2", Name: "imam"},
	}

	suite.useCaseMock.On("FindAllEmpl").Return(employee, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *EmployeeControllerSuite) TestListHandler_InternalServerError() {
	suite.useCaseMock.On("FindAllEmpl").Return([]model.Employee{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *EmployeeControllerSuite) TestGetHandler_Success() {
	employee := model.Employee{
		ID:   "1",
		Name: "imron",
	}

	suite.useCaseMock.On("FindByIdEmpl", "1").Return(employee, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *EmployeeControllerSuite) TestGetHandler_InternalServerError() {
	suite.useCaseMock.On("FindByIdEmpl", "1").Return(model.Employee{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

}

func (suite *EmployeeControllerSuite) TestUpdateHandler_StatusCreated() {
	payload := model.Employee{
		ID:   "1",
		Name: "imron",
	}

	reqBody, _ := json.Marshal(payload)

	suite.useCaseMock.On("UpdateEmpl", mock.AnythingOfType("model.Employee")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *EmployeeControllerSuite) TestUpdateHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *EmployeeControllerSuite) TestUpdateHandler_InternalServerError() {
	payload := model.Employee{
		Name: "imron",
	}

	suite.useCaseMock.On("UpdateEmpl", payload).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func TestEmployeeControllerSuite(t *testing.T) {
	suite.Run(t, new(EmployeeControllerSuite))
}
