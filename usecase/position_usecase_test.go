package usecase

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RepoPositionMock struct {
	mock.Mock
}

func (m *RepoPositionMock) Create(payload model.Position) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *RepoPositionMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *RepoPositionMock) Get(id string) (model.Position, error) {
	args := m.Called(id)
	return args.Get(0).(model.Position), args.Error(1)
}

func (m *RepoPositionMock) GetByName(name string) (model.Position, error) {
	args := m.Called(name)
	return args.Get(0).(model.Position), args.Error(1)
}

func (m *RepoPositionMock) List() ([]model.Position, error) {
	args := m.Called()
	return args.Get(0).([]model.Position), args.Error(1)
}

func (m *RepoPositionMock) Update(payload model.Position) error {
	args := m.Called(payload)
	return args.Error(0)
}

type PositionUseCaseSuite struct {
	suite.Suite
	repoMock *RepoPositionMock
	useCase  PositionUseCase
}

func (suite *PositionUseCaseSuite) SetupTest() {
	suite.repoMock = new(RepoPositionMock)
	suite.useCase = NewPositionUseCase(suite.repoMock)
}

var dataPositionDummy = []model.Position{
	{
		ID:   "1",
		Name: "Marketing",
	},
	{
		ID:   "2",
		Name: "HR",
	},
}

func (suite *PositionUseCaseSuite) TestRegisterNewPosition_Success() {
	payload := dataPositionDummy[0]
	suite.repoMock.On("GetByName", payload.Name).Return(model.Position{}, nil)
	suite.repoMock.On("Create", payload).Return(nil)

	err := suite.useCase.RegisterNewPosition(payload)
	assert.NoError(suite.T(), err)
}

func (suite *PositionUseCaseSuite) TestRegisterNewPosition_EmptyField() {
	payload := model.Position{}
	suite.repoMock.On("Create", payload).Return(fmt.Errorf("error"))
	err := suite.useCase.RegisterNewPosition(model.Position{})
	assert.Error(suite.T(), err)
}

func (suite *PositionUseCaseSuite) TestRegisterNewPosition_PositionExists() {
	payload := dataPositionDummy[0]
	existingPosition := dataPositionDummy[0]
	suite.repoMock.On("GetByName", payload.Name).Return(existingPosition, nil)

	err := suite.useCase.RegisterNewPosition(payload)
	expectedErrMsg := "position with name " + payload.Name + " exists"
	assert.EqualError(suite.T(), err, expectedErrMsg)
}

func (suite *PositionUseCaseSuite) TestRegisterNewPosition_Failed() {
	payload := model.Position{
		Name: "Marketing",
	}
	suite.repoMock.On("GetByName", payload.Name).Return(model.Position{}, nil)

	suite.repoMock.On("Create", payload).Return(fmt.Errorf("create error"))

	err := suite.useCase.RegisterNewPosition(payload)

	expectedErrMsg := "failed to create new position: create error"

	assert.EqualError(suite.T(), err, expectedErrMsg)
}

func (suite *PositionUseCaseSuite) TestFindByNamePosition_Success() {
	Name := "Annual Leave"
	expectedPosition := dataPositionDummy[0]
	suite.repoMock.On("GetByName", Name).Return(expectedPosition, nil)

	result, err := suite.useCase.GetByName(Name)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedPosition, result)

}

func (suite *PositionUseCaseSuite) TestFindByNamePosition_Failed() {
	Name := "Nonexistent Position"
	expectedError := fmt.Errorf("failed to get Position by name")
	suite.repoMock.On("GetByName", Name).Return(model.Position{}, expectedError)

	result, err := suite.useCase.GetByName(Name)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Empty(suite.T(), result)
}

func (suite *PositionUseCaseSuite) TestFindAllPosition_Success() {
	expectedPosition := dataPositionDummy

	suite.repoMock.On("List").Return(expectedPosition, nil)

	result, err := suite.useCase.FindAllPosition()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedPosition, result)

}

func (suite *PositionUseCaseSuite) TestFindAllPosition_Failed() {
	expectedError := fmt.Errorf("failed to get all Position")
	suite.repoMock.On("List").Return([]model.Position{}, expectedError)

	result, err := suite.useCase.FindAllPosition()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Empty(suite.T(), result)
}

func (suite *PositionUseCaseSuite) TestFindByIdPosition_Success() {
	positionID := "1"
	expectedPosition := model.Position{
		ID:   positionID,
		Name: "Marketing",
	}
	suite.repoMock.On("Get", positionID).Return(expectedPosition, nil)

	result, err := suite.useCase.FindByIdPosition(positionID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedPosition, result)

}

func (suite *PositionUseCaseSuite) TestFindByIdPosition_Failed() {
	positionID := "1"
	expectedError := fmt.Errorf("failed to get Position by ID")
	suite.repoMock.On("Get", positionID).Return(model.Position{}, expectedError)

	result, err := suite.useCase.FindByIdPosition(positionID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Equal(suite.T(), model.Position{}, result)
}

func (suite *PositionUseCaseSuite) TestUpdatePosition_Success() {
	payload := dataPositionDummy[0]
	suite.repoMock.On("GetByName", payload.Name).Return(model.Position{}, nil)
	suite.repoMock.On("Update", payload).Return(nil)

	err := suite.useCase.UpdatePosition(payload)
	assert.NoError(suite.T(), err)

}

func (suite *PositionUseCaseSuite) TestUpdatePosition_EmptyField() {
	payload := model.Position{
		Name: "",
	}

	suite.repoMock.On("GetByName", payload.Name).Return(model.Position{}, nil)
	suite.repoMock.On("Update", payload).Return(fmt.Errorf("name required field"))

	err := suite.useCase.UpdatePosition(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "name required field", err.Error())

}

func (suite *PositionUseCaseSuite) TestUpdatePosition_Failed() {
	payload := dataPositionDummy[0]

	suite.repoMock.On("GetByName", payload.Name).Return(model.Position{}, nil)
	suite.repoMock.On("Update", payload).Return(fmt.Errorf("update error"))

	err := suite.useCase.UpdatePosition(payload)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update error", err.Error())

}

func (suite *PositionUseCaseSuite) TestDeletePosition_Success() {
	positionID := "1"
	expectedPosition := model.Position{
		ID:   positionID,
		Name: "Marketing",
	}
	suite.repoMock.On("Get", positionID).Return(expectedPosition, nil)
	suite.repoMock.On("Delete", positionID).Return(nil)

	err := suite.useCase.DeletePosition(positionID)
	assert.NoError(suite.T(), err)

}

func (suite *PositionUseCaseSuite) TestDeletePosition_NotFound() {
	positionID := "1"

	suite.repoMock.On("Get", positionID).Return(model.Position{}, nil)
	suite.repoMock.On("Delete", positionID).Return(fmt.Errorf("not found"))

	err := suite.useCase.DeletePosition(positionID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not found", err.Error())

}

func (suite *PositionUseCaseSuite) TestDeletePosition_Failed() {
	positionID := "1"
	leavetype := dataPositionDummy[0]

	suite.repoMock.On("Get", positionID).Return(leavetype, nil)
	suite.repoMock.On("Delete", positionID).Return(fmt.Errorf("delete error"))

	err := suite.useCase.DeletePosition(positionID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete error", err.Error())

}

func TestPositionUseCaseSuite(t *testing.T) {
	suite.Run(t, new(PositionUseCaseSuite))
}
