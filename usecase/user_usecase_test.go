package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dataDummy = []model.UserCredential{
	{
		ID:       "1",
		Username: "agung",
		Password: "123",
		RoleId:   "",
		IsActive: true,
		Role: model.Role{
			Id:       "",
			RoleName: "",
		},
	},
	{
		ID:       "2",
		Username: "panji",
		Password: "123",
		RoleId:   "",
		IsActive: true,
		Role: model.Role{
			Id:       "",
			RoleName: "",
		},
	},
}

type RepoMock struct {
	mock.Mock
}

func (m *RepoMock) Create(payload model.UserCredential) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *RepoMock) Get(id string) (model.UserCredential, error) {
	args := m.Called(id)
	return args.Get(0).(model.UserCredential), args.Error(1)
}

func (m *RepoMock) GetByUsername(username string) (model.UserCredential, error) {
	args := m.Called(username)
	return args.Get(0).(model.UserCredential), args.Error(1)
}

func (m *RepoMock) GetByUsernamePassword(username string, password string) (model.UserCredential, error) {
	args := m.Called(username, password)
	return args.Get(0).(model.UserCredential), args.Error(1)
}

func (m *RepoMock) Paging(requestPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error) {
	args := m.Called(requestPaging)
	return args.Get(0).([]model.UserCredential), args.Get(1).(dto.Paging), args.Error(2)
}

// Update implements repository.StatusLeaveRepository.
func (m *RepoMock) Update(payload model.UserCredential) error {
	args := m.Called(payload)
	return args.Error(0)
}

type UserUseCaseSuite struct {
	suite.Suite
	repoMock *RepoMock
	useCase  UserUseCase
}

func (suite *UserUseCaseSuite) SetupTest() {
	suite.repoMock = new(RepoMock)
	suite.useCase = NewUserUseCase(suite.repoMock)
}

func (suite *UserUseCaseSuite) TestRegisterNewUser_Success() {
	dummy := dataDummy[0]
	suite.repoMock.On("Create", dummy).Return(nil)
	err := suite.useCase.RegisterNewUser(dummy)
	assert.Nil(suite.T(), err)
}

func (suite *UserUseCaseSuite) TestUpdateUser_Success() {
	dummy := dataDummy[0]
	dummy.Username = "anto"
	suite.repoMock.On("Update", dummy).Return(nil)
	err := suite.useCase.UpdateUser(dummy)
	assert.Nil(suite.T(), err)
}

func (suite *UserUseCaseSuite) TestFindByUsername_Success() {
	username := "agung"
	expectedUsername := dataDummy[0]
	suite.repoMock.On("GetByUsername", username).Return(expectedUsername, nil)

	result, err := suite.useCase.FindByUsername(username)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUsername, result)
}

func (suite *UserUseCaseSuite) TestFindByIdUser_Success() {
	userID := "1"
	expectedStatus := dataDummy[0]
	suite.repoMock.On("Get", userID).Return(expectedStatus, nil)

	result, err := suite.useCase.FindByIdUser(userID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatus, result)
}

func (suite *UserUseCaseSuite) TestFindByUsernamePassword_Success() {
	expectedUser := dataDummy[0]
	mockUsername := "agung"
	mockPassword := "password"

	suite.repoMock.On("GetByUsernamePassword", mockUsername, mockPassword).Return(expectedUser, nil)

	resultUser, err := suite.useCase.FindByUsernamePassword(mockUsername, mockPassword)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, resultUser)
}

func (suite *UserUseCaseSuite) TestFindAllEmpl_Success() {
	expectedData := []model.UserCredential{
		{
			ID:       "1",
			Username: "agung",
			Password: "password",
			RoleId:   "",
			IsActive: true,
			Role: model.Role{
				Id:       "",
				RoleName: "",
			},
		},
		{
			ID:       "2",
			Username: "panji",
			Password: "password",
			RoleId:   "",
			IsActive: true,
			Role: model.Role{
				Id:       "",
				RoleName: "",
			},
		},
	}
	expectedPaging := dto.Paging{
		Page:        1,
		RowsPerPage: 10,
		TotalRows:   len(expectedData),
		TotalPages:  1,
	}

	mockParam := dto.PaginationParam{Page: 1, Limit: 10}
	suite.repoMock.On("Paging", mockParam).Return(expectedData, expectedPaging, nil)

	resultData, resultPaging, err := suite.useCase.FindAllUser(mockParam)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedData, resultData)
	assert.Equal(suite.T(), expectedPaging, resultPaging)
}

func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseSuite))
}
