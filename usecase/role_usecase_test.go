package usecase

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type repoMock struct {
	mock.Mock
}

func (r *repoMock) Create(payload model.Role) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *repoMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *repoMock) Get(id string) (model.Role, error) {
	args := r.Called(id)
	return args.Get(0).(model.Role), args.Error(1)
}

func (r *repoMock) GetByName(roleName string) (model.Role, error) {
	args := r.Called(roleName)
	return args.Get(0).(model.Role), args.Error(1)
}

func (r *repoMock) List() ([]model.Role, error) {
	args := r.Called()
	return args.Get(0).([]model.Role), args.Error(1)
}

func (r *repoMock) Update(payload model.Role) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

type RoleUseCaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
	usecase  RoleUseCase
}

func (suite *RoleUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
	suite.usecase = NewRoleUseCase(suite.repoMock)
}

// Test Case
var roleDummy = []model.Role{
	{
		Id:              "1",
		RoleName: "Pending",
	},
	{
		Id:              "2",
		RoleName: "Approved",
	},
	{
		Id:              "3",
		RoleName: "Rejected",
	},
}

func (suite *RoleUseCaseTestSuite) TestRegisterNewRole_Success() {
	dmRole := roleDummy[0]

	// Set up the expectation for GetByName method
	suite.repoMock.On("GetByName", dmRole.RoleName).Return(model.Role{}, fmt.Errorf("not found"))

	// Set up the expectation for Create method
	suite.repoMock.On("Create", dmRole).Return(nil, fmt.Errorf("required fields"))

	err := suite.usecase.RegisterNewRole(dmRole)
	assert.Nil(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestRegisterNewRole_Fail() {
	dmRole := roleDummy[0]

	// Set up the expectation for GetByName method
	suite.repoMock.On("GetByName", dmRole.RoleName).Return(dmRole, nil)

	err := suite.usecase.RegisterNewRole(dmRole)
	assert.Error(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestFindByIdRole_Success() {
	dummy := roleDummy[0]
	suite.repoMock.On("Get", dummy.Id).Return(dummy, nil)
	actualRole, actualError := suite.usecase.FindByIdRole(dummy.Id)
	assert.Equal(suite.T(), dummy, actualRole)
	assert.Nil(suite.T(), actualError)
}

func (suite *RoleUseCaseTestSuite) TestFindByNameRole_Success() {
	dummy := roleDummy[0]

	// Set up the expectation for the GetByName method call
	suite.repoMock.On("GetByName", dummy.RoleName).Return(dummy, nil)

	actualRole, actualError := suite.usecase.FindByRolename(dummy.RoleName)
	assert.Equal(suite.T(), dummy, actualRole)
	assert.Nil(suite.T(), actualError)
}

func (suite *RoleUseCaseTestSuite) TestFindByIdRole_Fail() {
	suite.repoMock.On("Get", "1xxx").Return(model.Role{}, fmt.Errorf("error"))
	actualRole, actualError := suite.usecase.FindByIdRole("1xxx")
	assert.Equal(suite.T(), model.Role{}, actualRole)
	assert.Error(suite.T(), actualError)
}

func (suite *RoleUseCaseTestSuite) TestFindByAllRole_Success() {
	roles := roleDummy

	suite.repoMock.On("List").Return(roles, nil)
	foundRoles, err := suite.usecase.FindAllRole()

	assert.Equal(suite.T(), roles, foundRoles)
	assert.Nil(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestFindByAllRole_Fail() {
	suite.repoMock.On("List").Return([]model.Role{}, fmt.Errorf("error fetching roles"))
	foundRoles, err := suite.usecase.FindAllRole()

	assert.Empty(suite.T(), foundRoles)
	assert.Error(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestDeleteRole_Success() {
	dmRole := roleDummy[0]

	// Set up the expectation for Get method
	suite.repoMock.On("Get", dmRole.Id).Return(dmRole, nil)

	// Set up the expectation for Delete method
	suite.repoMock.On("Delete", dmRole.Id).Return(nil)

	err := suite.usecase.DeleteRole(dmRole.Id)
	assert.Nil(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestDeleteRole_Fail() {
	dmRole := roleDummy[0]

	// Set up the expectation for Get method
	suite.repoMock.On("Get", dmRole.Id).Return(dmRole, nil)

	// Set up the expectation for Delete method
	suite.repoMock.On("Delete", dmRole.Id).Return(fmt.Errorf("delete error"))

	err := suite.usecase.DeleteRole(dmRole.Id)
	assert.Error(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestUpdateRole_Success() {
	dmRole := roleDummy[0]

	// Set up the expectation for GetByName method
	suite.repoMock.On("GetByName", dmRole.RoleName).Return(dmRole, nil)

	// Set up the expectation for Update method
	suite.repoMock.On("Update", dmRole).Return(nil)

	err := suite.usecase.UpdateRole(dmRole)
	assert.Nil(suite.T(), err)
}

func (suite *RoleUseCaseTestSuite) TestUpdateRole_Fail() {
	dmRole := roleDummy[0]

	// Set up the expectation for GetByName method
	suite.repoMock.On("GetByName", dmRole.RoleName).Return(dmRole, nil)

	// Set up the expectation for Update method
	suite.repoMock.On("Update", dmRole).Return(fmt.Errorf("update error"))

	err := suite.usecase.UpdateRole(dmRole)
	assert.Error(suite.T(), err)
}

func TestRoleUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoleUseCaseTestSuite))
}
