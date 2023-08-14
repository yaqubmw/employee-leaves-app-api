package controller

import (
	"bytes"
	"employeeleave/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type roleUseCaseMock struct {
	mock.Mock
}

func (r *roleUseCaseMock) FindAllRole() ([]model.Role, error) {
	args := r.Called()
	return args.Get(0).([]model.Role), args.Error(1)
}

func (r *roleUseCaseMock) FindByRolename(roleName string) (model.Role, error) {
	args := r.Called(roleName)
	return args.Get(0).(model.Role), args.Error(1)
}

func (r *roleUseCaseMock) FindByIdRole(id string) (model.Role, error) {
	args := r.Called(id)
	return args.Get(0).(model.Role), args.Error(1)
}

func (r *roleUseCaseMock) RegisterNewRole(payload model.Role) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil 
}

func (r *roleUseCaseMock) UpdateRole(payload model.Role) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *roleUseCaseMock) DeleteRole(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type RoleControllerTestSuite struct {
	suite.Suite
	usecaseMock *roleUseCaseMock
	router *gin.Engine
	controller  *RoleController
	recorder *httptest.ResponseRecorder
}

func (suite *RoleControllerTestSuite) SetupTest() {
	suite.usecaseMock = new(roleUseCaseMock)
	suite.router = gin.Default()
	suite.recorder = httptest.NewRecorder()
	suite.controller = NewRoleController(suite.router, suite.usecaseMock)
}

func TestRoleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoleControllerTestSuite))
}

func (suite *RoleControllerTestSuite) TestCreateHandlerRole_Success() {
	dummyRequest := model.Role{
		RoleName: "roleName",
	}

	var newRole model.Role
	dummyRequest.Id = "1ABC"
	newRole.Id = dummyRequest.Id
	newRole.RoleName = dummyRequest.RoleName

	suite.usecaseMock.On("RegisterNewRole", newRole).Return(nil)
	payload, _ := json.Marshal(dummyRequest)

	request, _ := http.NewRequest(http.MethodPost, "/api/v1/roles", bytes.NewBuffer(payload))

	suite.router.ServeHTTP(suite.recorder, request)

	response := suite.recorder.Body.Bytes()
	actualRole := model.Role{}
	json.Unmarshal(response, &actualRole)
	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
	assert.Equal(suite.T(), dummyRequest, actualRole)
}

func (suite *RoleControllerTestSuite) TestCreateHandler_InternalServerError() {
	dummyRequest := model.Role{
		RoleName: "roleName",
	}

	reqBody, _ := json.Marshal(dummyRequest)
	suite.usecaseMock.On("RegisterNewRole", mock.AnythingOfType("model.Role")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/roles", bytes.NewBuffer(reqBody))
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


func (suite *RoleControllerTestSuite) TestCreateHandlerRole_BindingError() {
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/roles", nil)
	suite.router.ServeHTTP(suite.recorder, request)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *RoleControllerTestSuite) TestGetHandler_Success() {
	recorder := httptest.NewRecorder()
	role := model.Role{
		Id:              "1",
		RoleName: "roleName",
	}

	suite.usecaseMock.On("FindByIdRole", "1").Return(role, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/roles/1", nil)
	suite.router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

}

func (suite *RoleControllerTestSuite) TestGetHandler_InternalServerError() {
	suite.usecaseMock.On("FindByIdRole", "1").Return(model.Role{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/roles/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)

}

func (suite *RoleControllerTestSuite) TestUpdateHandler_StatusCreated() {
	payload := model.Role{
		Id:              "1",
		RoleName: "roleName",
	}

	reqBody, _ := json.Marshal(payload)

	suite.usecaseMock.On("UpdateRole", mock.AnythingOfType("model.Role")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/roles", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}


func (suite *RoleControllerTestSuite) TestListHandler_Success() {
	roles := []model.Role{
		{Id: "1", RoleName: "roleName1"},
		{Id: "2", RoleName: "roleName2"},
	}

	suite.usecaseMock.On("FindAllRole").Return(roles, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/roles", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *RoleControllerTestSuite) TestListHandler_InternalServerError() {
	suite.usecaseMock.On("FindAllRole").Return([]model.Role{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/roles", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *RoleControllerTestSuite) TestUpdateHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/roles", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *RoleControllerTestSuite) TestUpdateHandler_InternalServerError() {
	payload := model.Role{
		RoleName: "roleName",
	}

	suite.usecaseMock.On("UpdateRole", payload).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/roles", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}


func (suite *RoleControllerTestSuite) TestDeleteHandler_Success() {
	suite.usecaseMock.On("DeleteRole", "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/roles/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusNoContent, suite.recorder.Code)
	assert.Empty(suite.T(), suite.recorder.Body.String())
}

func (suite *RoleControllerTestSuite) TestDeleteHandler_InternalServerError() {
	suite.usecaseMock.On("DeleteRole", "1").Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/roles/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}
