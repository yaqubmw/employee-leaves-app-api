package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type historyRepoMock struct {
	mock.Mock
}

func (r *historyRepoMock) Create(payload model.HistoryLeave) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *historyRepoMock) Paging(requestPaging dto.PaginationParam) ([]model.HistoryLeave, dto.Paging, error) {
	args := r.Called(requestPaging)
	return args.Get(0).([]model.HistoryLeave), args.Get(1).(dto.Paging), args.Error(2)
}

func (r *historyRepoMock) GetHistoryById(id string) (model.HistoryLeave, error) {
	args := r.Called(id)
	return args.Get(0).(model.HistoryLeave), args.Error(1)
}

type HistoryUseCaseTestSuite struct {
	suite.Suite
	repoMock historyRepoMock
	usecase  HistoryUseCase
}

func (suite *HistoryUseCaseTestSuite) SetupTest() {
	suite.repoMock = historyRepoMock{}
	suite.usecase = NewHistoryUseCase(&suite.repoMock)
}

func (suite *HistoryUseCaseTestSuite) TestRegisterNewHistory_Success() {
	history := model.HistoryLeave{Id: "1", TransactionLeaveId: "123", DateEvent: time.Now()}

	suite.repoMock.On("Create", history).Return(nil)
	err := suite.usecase.RegisterNewHistory(history)
	assert.Nil(suite.T(), err)
}

func (suite *HistoryUseCaseTestSuite) TestRegisterNewHistory_Fail() {
	history := model.HistoryLeave{Id: "1", TransactionLeaveId: "123", DateEvent: time.Now()}

	suite.repoMock.On("Create", history).Return(fmt.Errorf("failed to create history"))
	err := suite.usecase.RegisterNewHistory(history)
	assert.Error(suite.T(), err)
}


func (suite *HistoryUseCaseTestSuite) TestFindHistoryById_Success() {
	historyID := "1"
	history := model.HistoryLeave{Id: "1", TransactionLeaveId: "123", DateEvent: time.Now()}

	suite.repoMock.On("GetHistoryById", historyID).Return(history, nil)
	foundHistory, err := suite.usecase.FindHistoryById(historyID)

	assert.Equal(suite.T(), history, foundHistory)
	assert.Nil(suite.T(), err)
}

func (suite *HistoryUseCaseTestSuite) TestFindHistoryById_Fail() {
	historyID := "1"

	suite.repoMock.On("GetHistoryById", historyID).Return(model.HistoryLeave{}, fmt.Errorf("not found"))
	foundHistory, err := suite.usecase.FindHistoryById(historyID)

	assert.Equal(suite.T(), model.HistoryLeave{}, foundHistory)
	assert.Error(suite.T(), err)
}

func (suite *HistoryUseCaseTestSuite) TestFindAllHistory_Success() {
	pagingParam := dto.PaginationParam{Page: 1, Limit: 10}
	historyList := []model.HistoryLeave{
		{Id: "1", TransactionLeaveId: "123", DateEvent: time.Now()},
		{Id: "2", TransactionLeaveId: "456", DateEvent: time.Now()},
	}

	pagingResponse := dto.Paging{
		Page:        1,
		RowsPerPage: 10,
		TotalRows:   len(historyList),
		TotalPages:  1,
	}

	suite.repoMock.On("Paging", pagingParam).Return(historyList, pagingResponse, nil)
	histories, paging, err := suite.usecase.FindAllHistory(pagingParam)

	assert.Equal(suite.T(), historyList, histories)
	assert.Equal(suite.T(), pagingResponse, paging)
	assert.Nil(suite.T(), err)
}

func (suite *HistoryUseCaseTestSuite) TestFindAllHistory_Fail() {
	pagingParam := dto.PaginationParam{Page: 1, Limit: 10}

	suite.repoMock.On("Paging", pagingParam).Return([]model.HistoryLeave{}, dto.Paging{}, fmt.Errorf("error fetching history"))
	histories, paging, err := suite.usecase.FindAllHistory(pagingParam)

	assert.Empty(suite.T(), histories)
	assert.Empty(suite.T(), paging)
	assert.Error(suite.T(), err)
}

func TestHistoryUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(HistoryUseCaseTestSuite))
}
