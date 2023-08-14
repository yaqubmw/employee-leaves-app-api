package usecase

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RepoLeaveTypeMock struct {
	mock.Mock
}

func (m *RepoLeaveTypeMock) Create(payload model.LeaveType) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *RepoLeaveTypeMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *RepoLeaveTypeMock) Get(id string) (model.LeaveType, error) {
	args := m.Called(id)
	return args.Get(0).(model.LeaveType), args.Error(1)
}

func (m *RepoLeaveTypeMock) GetByName(name string) (model.LeaveType, error) {
	args := m.Called(name)
	return args.Get(0).(model.LeaveType), args.Error(1)
}

func (m *RepoLeaveTypeMock) List() ([]model.LeaveType, error) {
	args := m.Called()
	return args.Get(0).([]model.LeaveType), args.Error(1)
}

func (m *RepoLeaveTypeMock) Update(payload model.LeaveType) error {
	args := m.Called(payload)
	return args.Error(0)
}

type LeaveTypeUseCaseSuite struct {
	suite.Suite
	repoMock *RepoLeaveTypeMock
	useCase  LeaveTypeUseCase
}

func (suite *LeaveTypeUseCaseSuite) SetupTest() {
	suite.repoMock = new(RepoLeaveTypeMock)
	suite.useCase = NewLeaveTypeUseCase(suite.repoMock)
}

var dataLeaveTypeDummy = []model.LeaveType{
	{
		ID:            "1",
		LeaveTypeName: "Matternity",
		QuotaLeave:    84,
	},
	{
		ID:            "2",
		LeaveTypeName: "Annual Leave",
		QuotaLeave:    12,
	},
}

func (suite *LeaveTypeUseCaseSuite) TestRegisterNewLeaveType_Success() {
	payload := dataLeaveTypeDummy[0]
	suite.repoMock.On("GetByName", payload.LeaveTypeName).Return(model.LeaveType{}, nil)
	suite.repoMock.On("Create", payload).Return(nil)

	err := suite.useCase.RegisterNewLeaveType(payload)
	assert.NoError(suite.T(), err)
}

func (suite *LeaveTypeUseCaseSuite) TestRegisterNewLeaveType_EmptyField() {
	payload := model.LeaveType{}
	suite.repoMock.On("Create", payload).Return(fmt.Errorf("error"))
	err := suite.useCase.RegisterNewLeaveType(model.LeaveType{})
	assert.Error(suite.T(), err)
}

func (suite *LeaveTypeUseCaseSuite) TestRegisterNewLeaveType_TypeExists() {
	payload := dataLeaveTypeDummy[0]
	existingType := dataLeaveTypeDummy[0]
	suite.repoMock.On("GetByName", payload.LeaveTypeName).Return(existingType, nil)

	err := suite.useCase.RegisterNewLeaveType(payload)
	expectedErrMsg := "leave type with name " + payload.LeaveTypeName + " exists"
	assert.EqualError(suite.T(), err, expectedErrMsg)
}

func (suite *LeaveTypeUseCaseSuite) TestRegisterNewLeaveType_Failed() {
	payload := model.LeaveType{
		LeaveTypeName: "Annual Leave",
	}
	suite.repoMock.On("GetByName", payload.LeaveTypeName).Return(model.LeaveType{}, nil)

	suite.repoMock.On("Create", payload).Return(fmt.Errorf("create error"))

	err := suite.useCase.RegisterNewLeaveType(payload)

	expectedErrMsg := "failed to create new leave type: create error"

	assert.EqualError(suite.T(), err, expectedErrMsg)
}

func (suite *LeaveTypeUseCaseSuite) TestFindByNameLeaveType_Success() {
	LeaveTypeName := "Annual Leave"
	expectedLeaveType := dataLeaveTypeDummy[0]
	suite.repoMock.On("GetByName", LeaveTypeName).Return(expectedLeaveType, nil)

	result, err := suite.useCase.GetByName(LeaveTypeName)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLeaveType, result)

}

func (suite *LeaveTypeUseCaseSuite) TestFindByNameLeaveType_Failed() {
	LeaveTypeName := "Nonexistent Leave Type"
	expectedError := fmt.Errorf("failed to get LeaveType by name")
	suite.repoMock.On("GetByName", LeaveTypeName).Return(model.LeaveType{}, expectedError)

	result, err := suite.useCase.GetByName(LeaveTypeName)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Empty(suite.T(), result)
}

func (suite *LeaveTypeUseCaseSuite) TestFindAllLeaveType_Success() {
	expectedLeaveTypes := dataLeaveTypeDummy

	suite.repoMock.On("List").Return(expectedLeaveTypes, nil)

	result, err := suite.useCase.FindAllLeaveType()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLeaveTypes, result)

}

func (suite *LeaveTypeUseCaseSuite) TestFindAllLeaveType_Failed() {
	expectedError := fmt.Errorf("failed to get all LeaveTypes")
	suite.repoMock.On("List").Return([]model.LeaveType{}, expectedError)

	result, err := suite.useCase.FindAllLeaveType()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Empty(suite.T(), result)
}

func (suite *LeaveTypeUseCaseSuite) TestFindByIdLeaveType_Success() {
	leaveTypeID := "1"
	expectedLeaveType := model.LeaveType{
		ID:            leaveTypeID,
		LeaveTypeName: "Annual Leave",
	}
	suite.repoMock.On("Get", leaveTypeID).Return(expectedLeaveType, nil)

	result, err := suite.useCase.FindByIdLeaveType(leaveTypeID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLeaveType, result)

}

func (suite *LeaveTypeUseCaseSuite) TestFindByIdLeaveType_Failed() {
	leaveTypeID := "1"
	expectedError := fmt.Errorf("failed to get LeaveType by ID")
	suite.repoMock.On("Get", leaveTypeID).Return(model.LeaveType{}, expectedError)

	result, err := suite.useCase.FindByIdLeaveType(leaveTypeID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Equal(suite.T(), model.LeaveType{}, result)
}

func (suite *LeaveTypeUseCaseSuite) TestUpdateLeaveType_Success() {
	payload := dataLeaveTypeDummy[0]
	suite.repoMock.On("GetByName", payload.LeaveTypeName).Return(model.LeaveType{}, nil)
	suite.repoMock.On("Update", payload).Return(nil)

	err := suite.useCase.UpdateLeaveType(payload)
	assert.NoError(suite.T(), err)

}

func (suite *LeaveTypeUseCaseSuite) TestUpdateLeaveType_EmptyField() {
	payload := model.LeaveType{
		LeaveTypeName: "",
	}

	suite.repoMock.On("GetByName", payload.LeaveTypeName).Return(model.LeaveType{}, nil)
	suite.repoMock.On("Update", payload).Return(fmt.Errorf("leave_type_name required field"))

	err := suite.useCase.UpdateLeaveType(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "leave_type_name required field", err.Error())

}

func (suite *LeaveTypeUseCaseSuite) TestUpdateLeaveType_Failed() {
	payload := dataLeaveTypeDummy[0]

	suite.repoMock.On("GetByName", payload.LeaveTypeName).Return(model.LeaveType{}, nil)
	suite.repoMock.On("Update", payload).Return(fmt.Errorf("update error"))

	err := suite.useCase.UpdateLeaveType(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update error", err.Error())

}

func (suite *LeaveTypeUseCaseSuite) TestDeleteLeaveType_Success() {
	leaveTypeID := "1"
	expectedLeaveType := model.LeaveType{
		ID:            leaveTypeID,
		LeaveTypeName: "Annual Leave",
	}
	suite.repoMock.On("Get", leaveTypeID).Return(expectedLeaveType, nil)
	suite.repoMock.On("Delete", leaveTypeID).Return(nil)

	err := suite.useCase.DeleteLeaveType(leaveTypeID)
	assert.NoError(suite.T(), err)

}

func (suite *LeaveTypeUseCaseSuite) TestDeleteLeaveType_NotFound() {
	leaveTypeID := "1"

	suite.repoMock.On("Get", leaveTypeID).Return(model.LeaveType{}, nil)
	suite.repoMock.On("Delete", leaveTypeID).Return(fmt.Errorf("not found"))

	err := suite.useCase.DeleteLeaveType(leaveTypeID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not found", err.Error())

}

func (suite *LeaveTypeUseCaseSuite) TestDeleteLeaveType_Failed() {
	leaveTypeID := "1"
	leavetype := dataLeaveTypeDummy[0]

	suite.repoMock.On("Get", leaveTypeID).Return(leavetype, nil)
	suite.repoMock.On("Delete", leaveTypeID).Return(fmt.Errorf("delete error"))

	err := suite.useCase.DeleteLeaveType(leaveTypeID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete error", err.Error())

}

func TestLeaveTypeUseCaseSuite(t *testing.T) {
	suite.Run(t, new(LeaveTypeUseCaseSuite))
}
