package usecase

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockStatusLeaveRepo struct {
	mock.Mock
}

// Create implements repository.StatusLeaveRepository.
func (m *MockStatusLeaveRepo) Create(payload model.StatusLeave) error {
	args := m.Called(payload)
	return args.Error(0)
}

// Delete implements repository.StatusLeaveRepository.
func (m *MockStatusLeaveRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Get implements repository.StatusLeaveRepository.
func (m *MockStatusLeaveRepo) Get(id string) (model.StatusLeave, error) {
	args := m.Called(id)
	return args.Get(0).(model.StatusLeave), args.Error(1)
}

// GetByNameStatus implements repository.StatusLeaveRepository.
func (m *MockStatusLeaveRepo) GetByNameStatus(statusLeaveName string) (model.StatusLeave, error) {
	args := m.Called(statusLeaveName)
	return args.Get(0).(model.StatusLeave), args.Error(1)
}

// List implements repository.StatusLeaveRepository.
func (m *MockStatusLeaveRepo) List() ([]model.StatusLeave, error) {
	args := m.Called()
	return args.Get(0).([]model.StatusLeave), args.Error(1)
}

// Update implements repository.StatusLeaveRepository.
func (m *MockStatusLeaveRepo) Update(payload model.StatusLeave) error {
	args := m.Called(payload)
	return args.Error(0)
}

type StatusLeaveUseCaseSuite struct {
	suite.Suite
	MockStatusLeaveRepo *MockStatusLeaveRepo
	useCase  StatusLeaveUseCase
}

func (suite *StatusLeaveUseCaseSuite) SetupTest() {
	suite.MockStatusLeaveRepo = new(MockStatusLeaveRepo)
	suite.useCase = NewStatusLeaveUseCase(suite.MockStatusLeaveRepo)
}

var statusDummy = []model.StatusLeave{
	{
		ID:              "1",
		StatusLeaveName: "Pending",
	},
	{
		ID:              "2",
		StatusLeaveName: "Approved",
	},
}

func (suite *StatusLeaveUseCaseSuite) TestRegisterNewStatusLeave_Success() {
	payload := statusDummy[0]
	suite.MockStatusLeaveRepo.On("GetByNameStatus", payload.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.MockStatusLeaveRepo.On("Create", payload).Return(nil)

	err := suite.useCase.RegisterNewStatusLeave(payload)
	assert.NoError(suite.T(), err)
}

func (suite *StatusLeaveUseCaseSuite) TestRegisterNewStatusLeave_EmptyField() {
	payload := model.StatusLeave{}
	suite.MockStatusLeaveRepo.On("Create", payload).Return(fmt.Errorf("error"))
	err := suite.useCase.RegisterNewStatusLeave(model.StatusLeave{})
	assert.Error(suite.T(), err)
}

func (suite *StatusLeaveUseCaseSuite) TestRegisterNewStatusLeave_StatusExists() {
	payload := model.StatusLeave{
		StatusLeaveName: "Pending",
	}
	existingStatus := statusDummy[0]

	suite.MockStatusLeaveRepo.On("GetByNameStatus", payload.StatusLeaveName).Return(existingStatus, nil)

	err := suite.useCase.RegisterNewStatusLeave(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "status with name Pending exists", err.Error())
}

func (suite *StatusLeaveUseCaseSuite) TestRegisterNewStatusLeave_CreateError() {
	payload := model.StatusLeave{
		StatusLeaveName: "Pending",
	}

	suite.MockStatusLeaveRepo.On("GetByNameStatus", payload.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.MockStatusLeaveRepo.On("Create", payload).Return(fmt.Errorf("create error"))

	err := suite.useCase.RegisterNewStatusLeave(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "failed to create new status: create error", err.Error())

}

func (suite *StatusLeaveUseCaseSuite) TestFindByNameStatusLeave_Success() {
	statusName := "Pending"
	expectedStatus := statusDummy[0]
	suite.MockStatusLeaveRepo.On("GetByNameStatus", statusName).Return(expectedStatus, nil)

	result, err := suite.useCase.FindByNameStatusLeave(statusName)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatus, result)

}

func (suite *StatusLeaveUseCaseSuite) TestFindAllStatusLeave_Success() {
	expectedStatuses := statusDummy

	suite.MockStatusLeaveRepo.On("List").Return(expectedStatuses, nil)

	result, err := suite.useCase.FindAllStatusLeave()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatuses, result)

}

func (suite *StatusLeaveUseCaseSuite) TestFindByIdStatusLeave_Success() {
	statusID := "1"
	expectedStatus := model.StatusLeave{
		ID:              statusID,
		StatusLeaveName: "Pending",
	}
	suite.MockStatusLeaveRepo.On("Get", statusID).Return(expectedStatus, nil)

	result, err := suite.useCase.FindByIdStatusLeave(statusID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatus, result)

}

func (suite *StatusLeaveUseCaseSuite) TestUpdateStatusLeave_Success() {
	payload := statusDummy[0]
	suite.MockStatusLeaveRepo.On("GetByNameStatus", payload.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.MockStatusLeaveRepo.On("Update", payload).Return(nil)

	err := suite.useCase.UpdateStatusLeave(payload)
	assert.NoError(suite.T(), err)

}

func (suite *StatusLeaveUseCaseSuite) TestUpdateStatusLeave_EmptyField() {
	payload := model.StatusLeave{
		StatusLeaveName: "",
	}

	err := suite.useCase.UpdateStatusLeave(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "status-leave-name required field", err.Error())

}

func (suite *StatusLeaveUseCaseSuite) TestUpdateStatusLeave_StatusExists() {
	payload := model.StatusLeave{
		StatusLeaveName: "Pending",
	}
	existingStatus := statusDummy[0]

	suite.MockStatusLeaveRepo.On("GetByNameStatus", payload.StatusLeaveName).Return(existingStatus, nil)

	err := suite.useCase.UpdateStatusLeave(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "status with name Pending exists", err.Error())

}

func (suite *StatusLeaveUseCaseSuite) TestUpdateStatusLeave_UpdateError() {
	payload := statusDummy[0]

	suite.MockStatusLeaveRepo.On("GetByNameStatus", payload.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.MockStatusLeaveRepo.On("Update", payload).Return(fmt.Errorf("update error"))

	err := suite.useCase.UpdateStatusLeave(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "failed to update status: update error", err.Error())

}

func (suite *StatusLeaveUseCaseSuite) TestDeleteStatusLeave_Success() {
	statusID := "1"
	expectedStatus := model.StatusLeave{
		ID:              statusID,
		StatusLeaveName: "Pending",
	}
	suite.MockStatusLeaveRepo.On("Get", statusID).Return(expectedStatus, nil)
	suite.MockStatusLeaveRepo.On("Delete", statusID).Return(nil)

	err := suite.useCase.DeleteStatusLeave(statusID)
	assert.NoError(suite.T(), err)

}

func (suite *StatusLeaveUseCaseSuite) TestDeleteStatusLeave_NotFound() {
	statusID := "1"

	suite.MockStatusLeaveRepo.On("Get", statusID).Return(model.StatusLeave{}, fmt.Errorf("not found"))

	err := suite.useCase.DeleteStatusLeave(statusID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "data with ID 1 not found", err.Error())

}

func (suite *StatusLeaveUseCaseSuite) TestDeleteStatusLeave_DeleteError() {
	statusID := "1"
	status := statusDummy[0]

	suite.MockStatusLeaveRepo.On("Get", statusID).Return(status, nil)
	suite.MockStatusLeaveRepo.On("Delete", statusID).Return(fmt.Errorf("delete error"))

	err := suite.useCase.DeleteStatusLeave(statusID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "failed to delete statusLeave: delete error", err.Error())

}

func TestStatusLeaveUseCaseSuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveUseCaseSuite))
}
