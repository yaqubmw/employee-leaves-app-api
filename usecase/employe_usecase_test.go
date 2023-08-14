package usecase

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmployeeRepoMock struct {
	mock.Mock
}

func (m *EmployeeRepoMock) Create(payload model.Employee) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *EmployeeRepoMock) Get(id string) (model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *EmployeeRepoMock) GetByName(name string) (model.Employee, error) {
	args := m.Called(name)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *EmployeeRepoMock) List() ([]model.Employee, error) {
	args := m.Called()
	return args.Get(0).([]model.Employee), args.Error(1)
}

func (m *EmployeeRepoMock) Update(payload model.Employee) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *EmployeeRepoMock) UpdateAnnualLeave(id string, availableDays int) error {
	args := m.Called(id, availableDays)
	return args.Error(0)
}

func (m *EmployeeRepoMock) UpdateMaternityLeave(id string, availableDays int) error {
	args := m.Called(id, availableDays)
	return args.Error(0)
}

func (m *EmployeeRepoMock) UpdateMarriageLeave(id string, availableDays int) error {
	args := m.Called(id, availableDays)
	return args.Error(0)
}

func (m *EmployeeRepoMock) UpdateMenstrualLeave(id string, availableDays int) error {
	args := m.Called(id, availableDays)
	return args.Error(0)
}

func (m *EmployeeRepoMock) UpdatePaternityLeave(id string, availableDays int) error {
	args := m.Called(id, availableDays)
	return args.Error(0)
}

type EmployeeUseCaseSuite struct {
	suite.Suite
	repoMock *EmployeeRepoMock
	useCase  EmployeeUseCase
}

func (suite *EmployeeUseCaseSuite) SetupTest() {
	suite.repoMock = new(EmployeeRepoMock)
	suite.useCase = NewEmplUseCase(suite.repoMock)
}

var employeeDummy = []model.Employee{
	{
		ID:          "1",
		Name:        "imron",
		PhoneNumber: "081281164811",
		Email:       "iyaron@gmail.com",
		Address:     "Jakarta",
	},
	{
		ID:          "2",
		Name:        "imam",
		PhoneNumber: "081281164811",
		Email:       "imam@gmail.com",
		Address:     "tanggerang",
	},
}

func (suite *EmployeeUseCaseSuite) TestRegisterNewEmpl_Success() {
	payload := employeeDummy[0]
	suite.repoMock.On("GetByName", payload.Name).Return(model.Employee{}, nil)
	suite.repoMock.On("Create", payload).Return(nil)

	err := suite.useCase.RegisterNewEmpl(payload)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestRegisterNewEmpl_EmptyField() {
	payload := model.Employee{}
	suite.repoMock.On("Create", payload).Return(fmt.Errorf("error"))

	err := suite.useCase.RegisterNewEmpl(payload)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestRegisterNewEmpl_EmployeeNameExists() {
	payload := model.Employee{
		Name: "imron", //name exists
	}
	existingEmployee := employeeDummy[0]

	suite.repoMock.On("GetByName", payload.Name).Return(existingEmployee, nil)

	err := suite.useCase.RegisterNewEmpl(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "employee with name imron exists", err.Error())
}

func (suite *EmployeeUseCaseSuite) TestRegisterNewEmpl_CreateError() {
	payload := model.Employee{
		ID:          "1",
		Name:        "imron",
		PhoneNumber: "081281164811",
		Email:       "iyaron@gmail.com",
		Address:     "Jakarta",
	}

	suite.repoMock.On("GetByName", payload.Name).Return(model.Employee{}, nil)
	suite.repoMock.On("Create", payload).Return(fmt.Errorf("create error"))

	err := suite.useCase.RegisterNewEmpl(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "failed to create new employee: create error", err.Error())
}

func (suite *EmployeeUseCaseSuite) TestFindAllEmpl_Success() {
	expectedEmployees := employeeDummy

	suite.repoMock.On("List").Return(expectedEmployees, nil)

	result, err := suite.useCase.FindAllEmpl()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedEmployees, result)
}

func (suite *EmployeeUseCaseSuite) TestFindByIdEmpl_Success() {
	expectedEmployee := model.Employee{
		ID:   "1",
		Name: "imron",
	}
	suite.repoMock.On("Get", "1").Return(expectedEmployee, nil)

	result, err := suite.useCase.FindByIdEmpl("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedEmployee, result)
}

func (suite *EmployeeUseCaseSuite) TestUpdateEmpl_Success() {
	payload := employeeDummy[0]
	suite.repoMock.On("Update", payload).Return(nil)

	err := suite.useCase.UpdateEmpl(payload)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateEmpl_UpdateError() {
	payload := employeeDummy[0]
	expectedError := fmt.Errorf("update error")
	suite.repoMock.On("Update", payload).Return(expectedError)

	err := suite.useCase.UpdateEmpl(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), fmt.Sprintf("failed to update employee: %v", expectedError), err.Error())
}

func (suite *EmployeeUseCaseSuite) TestPaternityLeave_Success() {
	id := "1"
	availableDays := 5
	suite.repoMock.On("PaternityLeave", id, availableDays).Return(nil)

	err := suite.useCase.PaternityLeave(id, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateAnnualLeave_Success() {
	id := "1"
	availableDays := 20
	suite.repoMock.On("UpdateAnnualLeave", id, availableDays).Return(nil)

	err := suite.useCase.UpdateAnnualLeave(id, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateMarriageLeave_UpdateError() {
	id := "1"
	availableDays := 7
	expectedError := fmt.Errorf("update error")
	suite.repoMock.On("UpdateMarriageLeave", id, availableDays).Return(expectedError)

	err := suite.useCase.UpdateMarriageLeave(id, availableDays)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateMaternityLeave_Success() {
	id := "1"
	availableDays := 30
	suite.repoMock.On("UpdateMaternityLeave", id, availableDays).Return(nil)

	err := suite.useCase.UpdateMaternityLeave(id, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateMenstrualLeave_Success() {
	id := "1"
	availableDays := 5
	suite.repoMock.On("UpdateMenstrualLeave", id, availableDays).Return(nil)

	err := suite.useCase.UpdateMenstrualLeave(id, availableDays)
	assert.NoError(suite.T(), err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateMaternityLeave_UpdateError() {
	id := "1"
	availableDays := 10
	expectedError := fmt.Errorf("update error")
	suite.repoMock.On("UpdateMaternityLeave", id, availableDays).Return(expectedError)

	err := suite.useCase.UpdateMaternityLeave(id, availableDays)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *EmployeeUseCaseSuite) TestUpdateMenstrualLeave_UpdateError() {
	id := "1"
	availableDays := 2
	expectedError := fmt.Errorf("update error")
	suite.repoMock.On("UpdateMenstrualLeave", id, availableDays).Return(expectedError)

	err := suite.useCase.UpdateMenstrualLeave(id, availableDays)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func TestEmployeeUseCaseSuite(t *testing.T) {
	suite.Run(t, new(EmployeeUseCaseSuite))
}
