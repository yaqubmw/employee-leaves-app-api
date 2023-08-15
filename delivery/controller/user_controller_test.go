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
	"employeeleave/model/dto"
)

type userUCMock struct {
	mock.Mock
}

func (u *userUCMock) FindAllUser(requesPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error) {
	args := u.Called()
	return args.Get(0).([]model.UserCredential), args.Get(0).(dto.Paging), args.Error(1)
}

func (u *userUCMock) FindByIdUser(id string) (model.UserCredential, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.UserCredential{}, args.Error(1)
	}
	return args.Get(0).(model.UserCredential), nil
}

func (s *userUCMock) FindByUsername(username string) (model.UserCredential, error) {
	args := s.Called(username)
	if args.Get(1) != nil {
		return model.UserCredential{}, args.Error(1)
	}
	return args.Get(0).(model.UserCredential), nil
}

func (s *userUCMock) FindByUsernamePassword(username, password string) (model.UserCredential, error) {
	args := s.Called(username, password)
	if args.Get(1) != nil {
		return model.UserCredential{}, args.Error(1)
	}
	return args.Get(0).(model.UserCredential), nil
}

func (u *userUCMock) RegisterNewUser(payload model.UserCredential) error {
	args := u.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (u *userUCMock) UpdateUser(payload model.UserCredential) error {
	args := u.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type UserControllerSuite struct {
	suite.Suite
	controller  *UserController
	useCaseMock *userUCMock
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
}

func (suite *UserControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.useCaseMock = &userUCMock{}
	suite.controller = NewUserController(suite.router, suite.useCaseMock)
	suite.recorder = httptest.NewRecorder()
}

func (suite *UserControllerSuite) TestCreateHandler_StatusCreated() {
	payload := model.UserCredential{
		ID:       "1",
		Username: "agung",
		Password: "password",
		RoleId:   "",
		IsActive: true,
		Role:     model.Role{},
	}

	suite.useCaseMock.On("RegisterNewUser", mock.AnythingOfType("model.UserCredential")).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *UserControllerSuite) TestCreateHandler_BadRequest() {
	payload := []model.UserCredential{}
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *UserControllerSuite) TestCreateHandler_InternalServerError() {
	payload := model.UserCredential{
		ID:       "",
		Username: "",
		Password: "",
		RoleId:   "",
		IsActive: false,
		Role:     model.Role{},
	}
	reqBody, _ := json.Marshal(payload)
	suite.useCaseMock.On("RegisterNewUser", mock.AnythingOfType("model.UserCredential")).Return(fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
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

// func (suite *UserControllerSuite) TestListHandler_Success() {
// 	users := []model.StatusLeave{
// 		{ID: "1", StatusLeaveName: "Pending"},
// 		{ID: "2", StatusLeaveName: "Approved"},
// 	}

// 	suite.useCaseMock.On("FindAllStatusLeave").Return(statusLeaves, nil)

// 	req, _ := http.NewRequest(http.MethodGet, "/api/v1/statusleaves", nil)
// 	suite.router.ServeHTTP(suite.recorder, req)

// 	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
// }

// func (suite *StatusLeaveControllerSuite) TestListHandler_InternalServerError() {
// 	suite.useCaseMock.On("FindAllStatusLeave").Return([]model.StatusLeave{}, fmt.Errorf("error"))

// 	req, _ := http.NewRequest(http.MethodGet, "/api/v1/statusleaves", nil)
// 	suite.router.ServeHTTP(suite.recorder, req)

// 	response := suite.recorder.Body.Bytes()
// 	var actualError struct {
// 		Err string `json:"err"`
// 	}
// 	json.Unmarshal(response, &actualError)

// 	assert.Equal(suite.T(), "error", actualError.Err)
// 	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
// }

func (suite *UserControllerSuite) TestGetHandler_Success() {
	users := model.UserCredential{
		ID:       "1",
		Username: "agung",
		Password: "password",
		RoleId:   "",
		IsActive: true,
		Role:     model.Role{},
	}

	suite.useCaseMock.On("FindByIdUser", "1").Return(users, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *UserControllerSuite) TestGetHandler_InternalServerError() {
	suite.useCaseMock.On("FindByIdUser", "1").Return(model.UserCredential{}, fmt.Errorf("error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	response := suite.recorder.Body.Bytes()
	var actualError struct {
		Err string `json:"err"`
	}
	json.Unmarshal(response, &actualError)

	assert.Equal(suite.T(), "error", actualError.Err)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func (suite *UserControllerSuite) TestUpdateHandler_StatusCreated() {
	payload := model.UserCredential{
		ID:       "1",
		Username: "panji",
		Password: "passw0rd",
		RoleId:   "",
		IsActive: true,
		Role:     model.Role{},
	}

	reqBody, _ := json.Marshal(payload)

	suite.useCaseMock.On("UpdateUser", mock.AnythingOfType("model.UserCredential")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}

func (suite *UserControllerSuite) TestUpdateHandler_BadRequest() {
	payload := []byte(`invalid json`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *UserControllerSuite) TestUpdateHandler_InternalServerError() {
	payload := model.UserCredential{
		ID:       "",
		Username: "",
		Password: "",
		RoleId:   "",
		IsActive: false,
		Role:     model.Role{},
	}

	suite.useCaseMock.On("UpdateUser", payload).Return(fmt.Errorf("error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(suite.recorder, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
