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

type PositionUCMock struct {
	mock.Mock
}

func (s *PositionUCMock) DeletePosition(id string) error {
	args := s.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *PositionUCMock) FindAllPosition() ([]model.Position, error) {
	args := s.Called()
	return args.Get(0).([]model.Position), args.Error(1)
}

func (s *PositionUCMock) FindByIdPosition(id string) (model.Position, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return model.Position{}, args.Error(1)
	}
	return args.Get(0).(model.Position), nil
}

func (s *PositionUCMock) GetByName(Name string) (model.Position, error) {
	args := s.Called(Name)
	if args.Get(1) != nil {
		return model.Position{}, args.Error(1)
	}
	return args.Get(0).(model.Position), nil
}

func (s *PositionUCMock) RegisterNewPosition(payload model.Position) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *PositionUCMock) UpdatePosition(payload model.Position) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type PositionControllerSuite struct {
	suite.Suite
	controller  *PositionController
	useCaseMock *PositionUCMock
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
}

func (suite *PositionControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.useCaseMock = &PositionUCMock{}
	suite.controller = NewPositionController(suite.useCaseMock, suite.router)
	suite.recorder = httptest.NewRecorder()
}

func (suite *PositionControllerSuite) TestCreateHandler_PositionCreated() {
	payload := model.Position{
		Name: "Marketing",
	}

	suite.useCaseMock.On("RegisterNewPosition", mock.AnythingOfType("model.Position")).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/positions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestCreateHandler_BadRequest() {
	payload := []model.Position{}
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/positions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestCreateHandler_InternalServerError() {
	payload := model.Position{
		Name: "Marketing",
	}
	reqBody, _ := json.Marshal(payload)
	suite.useCaseMock.On("RegisterNewPosition", mock.AnythingOfType("model.Position")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/positions", bytes.NewBuffer(reqBody))
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

func (suite *PositionControllerSuite) TestListPositionHandler_Success() {
	positions := []model.Position{
		{ID: "1", Name: "Marketing"},
		{ID: "2", Name: "Social Media"},
	}

	suite.useCaseMock.On("FindAllPosition").Return(positions, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/positions", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *PositionControllerSuite) TestListPositionHandler_InternalServerError() {
	suite.useCaseMock.On("FindAllPosition").Return([]model.Position{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/positions", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestGetPositionHandler_Success() {
	position := model.Position{
		ID:   "1",
		Name: "Marketing",
	}

	suite.useCaseMock.On("FindByIdPosition", "1").Return(position, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/positions/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

}

func (suite *PositionControllerSuite) TestGetPositionHandler_InternalServerError() {
	suite.useCaseMock.On("FindByIdPosition", "1").Return(model.Position{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/positions/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

}

func (suite *PositionControllerSuite) TestGetByNameHandler_Success() {
	expectedPosition := model.Position{
		Name: "Marketing",
	}

	suite.useCaseMock.On("GetByName", "Marketing").Return(expectedPosition, nil)

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/positions/name/Marketing", nil)
	suite.router.ServeHTTP(suite.recorder, request)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

	var response struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Data model.Position `json:"data"`
	}

	err := json.Unmarshal(suite.recorder.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), 200, response.Status.Code)
	assert.Equal(suite.T(), "Get By Name Data Successfully", response.Status.Description)
	assert.Equal(suite.T(), expectedPosition, response.Data)
}

func (suite *PositionControllerSuite) TestGetByNameHandler_InternalServerError() {
	suite.useCaseMock.On("GetByName", "Marketing").Return(model.Position{}, fmt.Errorf("internal server error"))

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/positions/name/Marketing", nil)
	suite.router.ServeHTTP(suite.recorder, request)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}

	err := json.Unmarshal(response, &actualError)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "internal server error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestUpdatePositionHandler_PositionCreated() {
	payload := model.Position{
		ID:   "1",
		Name: "Marekting",
	}

	reqBody, _ := json.Marshal(payload)

	suite.useCaseMock.On("UpdatePosition", mock.AnythingOfType("model.Position")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/positions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestUpdatePositionHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/positions", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestUpdatePositionHandler_InternalServerError() {
	payload := model.Position{
		Name: "Marketing",
	}

	suite.useCaseMock.On("UpdatePosition", payload).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/positions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *PositionControllerSuite) TestDeleteHandler_Success() {
	suite.useCaseMock.On("DeletePosition", "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/positions/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusNoContent, suite.recorder.Code)
	assert.Empty(suite.T(), suite.recorder.Body.String())
}

func (suite *PositionControllerSuite) TestDeleteHandler_InternalServerError() {
	suite.useCaseMock.On("DeletePosition", "1").Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/positions/id/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func TestPositionControllerSuite(t *testing.T) {
	suite.Run(t, new(PositionControllerSuite))
}
