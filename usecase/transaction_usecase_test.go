package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTransactionLeaveRepo struct {
	mock.Mock
}

func (m *MockTransactionLeaveRepo) Create(payload model.TransactionLeave) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *MockTransactionLeaveRepo) GetByEmployeeID(employeeID string) ([]model.TransactionLeave, error) {
	args := m.Called(employeeID)
	return args.Get(0).([]model.TransactionLeave), args.Error(1)
}

func (m *MockTransactionLeaveRepo) GetByID(id string) (model.TransactionLeave, error) {
	args := m.Called(id)
	return args.Get(0).(model.TransactionLeave), args.Error(1)
}

func (m *MockTransactionLeaveRepo) GetByIdTxNonDto(id string) (model.TransactionLeave, error) {
	args := m.Called(id)
	return args.Get(0).(model.TransactionLeave), args.Error(1)
}

func (m *MockTransactionLeaveRepo) Paging(requestPagung dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error) {
	args := m.Called(requestPagung)
	return args.Get(0).([]dto.TransactionResponseDto), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockTransactionLeaveRepo) UpdateStatus(transactionID string, statusID string) error {
	args := m.Called(transactionID, statusID)
	return args.Error(0)
}

type MockEmployeeUseCase struct {
	mock.Mock
}

func (mock *MockEmployeeUseCase) FindAllEmpl() ([]model.Employee, error) {
	args := mock.Called()
	return args.Get(0).([]model.Employee), args.Error(1)
}

func (mock *MockEmployeeUseCase) FindByIdEmpl(id string) (model.Employee, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (mock *MockEmployeeUseCase) PaternityLeave(id string, availableDays int) error {
	args := mock.Called(id, availableDays)
	return args.Error(0)
}

func (mock *MockEmployeeUseCase) RegisterNewEmpl(payload model.Employee) error {
	args := mock.Called(payload)
	return args.Error(0)
}

func (mock *MockEmployeeUseCase) UpdateAnnualLeave(id string, availableDays int) error {
	args := mock.Called(id, availableDays)
	return args.Error(0)
}

func (mock *MockEmployeeUseCase) UpdateEmpl(payload model.Employee) error {
	args := mock.Called(payload)
	return args.Error(0)
}

func (mock *MockEmployeeUseCase) UpdateMarriageLeave(id string, availableDays int) error {
	args := mock.Called(id, availableDays)
	return args.Error(0)
}

func (mock *MockEmployeeUseCase) UpdateMaternityLeave(id string, availableDays int) error {
	args := mock.Called(id, availableDays)
	return args.Error(0)
}

func (mock *MockEmployeeUseCase) UpdateMenstrualLeave(id string, availableDays int) error {
	args := mock.Called(id, availableDays)
	return args.Error(0)
}

type MockLeaveTypeUseCase struct {
	mock.Mock
}

func (mock *MockLeaveTypeUseCase) DeleteLeaveType(id string) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *MockLeaveTypeUseCase) FindAllLeaveType() ([]model.LeaveType, error) {
	args := mock.Called()
	return args.Get(0).([]model.LeaveType), args.Error(1)
}

func (mock *MockLeaveTypeUseCase) FindByIdLeaveType(id string) (model.LeaveType, error) {
	args := mock.Called(id)
	return args.Get(0).(model.LeaveType), args.Error(1)
}

func (mock *MockLeaveTypeUseCase) FindRoleNameId(id string) (model.Role, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Role), args.Error(1)
}

func (mock *MockLeaveTypeUseCase) GetByName(name string) (model.LeaveType, error) {
	args := mock.Called(name)
	return args.Get(0).(model.LeaveType), args.Error(1)
}

func (mock *MockLeaveTypeUseCase) RegisterNewLeaveType(payload model.LeaveType) error {
	args := mock.Called(payload)
	return args.Error(0)
}

func (mock *MockLeaveTypeUseCase) UpdateLeaveType(payload model.LeaveType) error {
	args := mock.Called(payload)
	return args.Error(0)
}

type MockStatusLeaveUseCase struct {
	mock.Mock
}

func (m *MockStatusLeaveUseCase) DeleteStatusLeave(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockStatusLeaveUseCase) FindAllStatusLeave() ([]model.StatusLeave, error) {
	args := m.Called()
	return args.Get(0).([]model.StatusLeave), args.Error(1)
}

func (m *MockStatusLeaveUseCase) FindByIdStatusLeave(id string) (model.StatusLeave, error) {
	args := m.Called(id)
	return args.Get(0).(model.StatusLeave), args.Error(1)
}

func (m *MockStatusLeaveUseCase) FindByNameStatusLeave(statusName string) (model.StatusLeave, error) {
	args := m.Called(statusName)
	return args.Get(0).(model.StatusLeave), args.Error(1)
}

func (m *MockStatusLeaveUseCase) RegisterNewStatusLeave(payload model.StatusLeave) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *MockStatusLeaveUseCase) UpdateStatusLeave(payload model.StatusLeave) error {
	args := m.Called(payload)
	return args.Error(0)
}

type TransactionLeaveUseCaseSuite struct {
	suite.Suite
	MockTransactionLeaveRepo *MockTransactionLeaveRepo
	employeeUC               *MockEmployeeUseCase
	leaveTypeUC              *MockLeaveTypeUseCase
	statusLeaveUC            *MockStatusLeaveUseCase
	transactionUC            TransactionLeaveUseCase
}

func (suite *TransactionLeaveUseCaseSuite) SetupTest() {
	suite.MockTransactionLeaveRepo = new(MockTransactionLeaveRepo)
	suite.employeeUC = new(MockEmployeeUseCase)
	suite.leaveTypeUC = new(MockLeaveTypeUseCase)
	suite.statusLeaveUC = new(MockStatusLeaveUseCase)

	suite.transactionUC = NewTransactionLeaveUseCase(suite.MockTransactionLeaveRepo, suite.employeeUC, suite.leaveTypeUC, suite.statusLeaveUC)
}

// code test

func (suite *TransactionLeaveUseCaseSuite) TestApplyLeave_FindEmployeeError() {

	trx := model.TransactionLeave{
		EmployeeID: "123",
	}
	expectedError := errors.New("employee not found")

	suite.employeeUC.On("FindByIdEmpl", trx.EmployeeID).Return(model.Employee{}, expectedError)

	err := suite.transactionUC.ApplyLeave(trx)

	// Assertions
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())

	suite.employeeUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApplyLeave_FindLeaveTypeByIdError() {

	trx := model.TransactionLeave{
		LeaveTypeID: "123",
	}
	expectedError := errors.New("leave type not found")
	suite.leaveTypeUC.On("FindByIdLeaveType", trx.LeaveTypeID).Return(model.LeaveType{}, expectedError)
	suite.employeeUC.On("FindByIdEmpl", trx.EmployeeID).Return(model.Employee{}, nil)

	err := suite.transactionUC.ApplyLeave(trx)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())

	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApplyLeave_FindStatusLeaveByNameError() {

	trx := model.TransactionLeave{
		StatusLeaveID: "123",
	}
	expectedError := errors.New("status not found")

	suite.statusLeaveUC.On("FindByNameStatusLeave", "Pending").Return(model.StatusLeave{}, expectedError)
	suite.employeeUC.On("FindByIdEmpl", trx.EmployeeID).Return(model.Employee{}, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", trx.LeaveTypeID).Return(model.LeaveType{}, nil)

	err := suite.transactionUC.ApplyLeave(trx)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())

	suite.statusLeaveUC.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
}
func (suite *TransactionLeaveUseCaseSuite) TestApplyLeave() {
	trx := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		// ... other fields
	}

	expectedEmployee := model.Employee{
		ID: "employee_id",
	}
	expectedLeaveType := model.LeaveType{
		ID: "leave_type_id",
	}
	expectedStatusLeave := model.StatusLeave{
		ID: "status_leave_id",
	}
	expectedHistoryLeaves := model.HistoryLeave{
		Id:                 "0efaeae0-f3b3-48ee-9a84-289b6f3015d6",
		TransactionLeaveId: trx.ID,
		DateEvent:          time.Now(),
	}

	suite.employeeUC.On("FindByIdEmpl", trx.EmployeeID).Return(expectedEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", trx.LeaveTypeID).Return(expectedLeaveType, nil)
	suite.statusLeaveUC.On("FindByNameStatusLeave", "Pending").Return(expectedStatusLeave, nil)

	suite.MockTransactionLeaveRepo.On("Create", mock.MatchedBy(func(input model.TransactionLeave) bool {
		return input.ID == trx.ID &&
			input.EmployeeID == expectedEmployee.ID &&
			input.LeaveTypeID == expectedLeaveType.ID &&
			input.StatusLeaveID == expectedStatusLeave.ID &&
			len(input.HistoryLeaves.Id) > 0 &&
			input.HistoryLeaves.TransactionLeaveId == trx.ID &&
			input.HistoryLeaves.DateEvent.Sub(expectedHistoryLeaves.DateEvent) <= time.Second
	})).Return(nil).Once()

	err := suite.transactionUC.ApplyLeave(trx)
	assert.NoError(suite.T(), err)

	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())

	suite.MockTransactionLeaveRepo.On("Create", mock.Anything).Return(fmt.Errorf("an error occurred")).Once()
	err = suite.transactionUC.ApplyLeave(trx)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to register new transaction an error occurred")
}

func (suite *TransactionLeaveUseCaseSuite) TestFindByIdTrx_Success() {
	expectedTrx := model.TransactionLeave{}
	suite.MockTransactionLeaveRepo.On("GetByID", "1").Return(expectedTrx, nil)

	result, err := suite.transactionUC.FindById("1")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedTrx, result)
}

func (suite *TransactionLeaveUseCaseSuite) TestFindByIdEmpl_Success() {
	expectedData := []model.TransactionLeave{
		{
			ID:         "1",
			EmployeeID: "123",
		},
	}
	suite.MockTransactionLeaveRepo.On("GetByEmployeeID", "123").Return(expectedData, nil)

	result, err := suite.transactionUC.FindByIdEmpl("123")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedData, result)
}

func (suite *TransactionLeaveUseCaseSuite) TestFindAllEmpl_Success() {
	expectedData := []dto.TransactionResponseDto{
		{
			ID:        "1",
			DateStart: "2023-08-01",
			DateEnd:   "2023-08-05",
			DayType:   "Full Day",
			Reason:    "Sick",
		},
		{
			ID:        "2",
			DateStart: "2023-08-06",
			DateEnd:   "2023-08-10",
			DayType:   "Half Day",
			Reason:    "Personal",
		},
	}
	expectedPaging := dto.Paging{
		Page:        1,
		RowsPerPage: 10,
		TotalRows:   len(expectedData),
		TotalPages:  1,
	}

	mockParam := dto.PaginationParam{Page: 1, Limit: 10}
	suite.MockTransactionLeaveRepo.On("Paging", mockParam).Return(expectedData, expectedPaging, nil)

	resultData, resultPaging, err := suite.transactionUC.FindAllEmpl(mockParam)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedData, resultData)
	assert.Equal(suite.T(), expectedPaging, resultPaging)
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_AnnualApproved() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:          "employee_id",
		AnnualLeave: 10,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "annual",
	}

	// suite.statusLeaveUC.On("FindByNameStatusLeave", mockStatusLeave.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil) // Updated this line

	// Correct the arguments for UpdateAnnualLeave mock
	suite.employeeUC.On("UpdateAnnualLeave", mockTransaction.EmployeeID, 7).Return(nil)

	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockStatusLeave.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MaternityApproved() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		MaternityLeave: 10,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "maternity",
	}

	// suite.statusLeaveUC.On("FindByNameStatusLeave", mockStatusLeave.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil) // Updated this line

	// Correct the arguments for UpdateAnnualLeave mock
	suite.employeeUC.On("UpdateMaternityLeave", mockTransaction.EmployeeID, 7).Return(nil)

	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockStatusLeave.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MarriageApproved() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:            "employee_id",
		MarriageLeave: 10,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "marriage",
	}

	// suite.statusLeaveUC.On("FindByNameStatusLeave", mockStatusLeave.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil) // Updated this line

	// Correct the arguments for UpdateAnnualLeave mock
	suite.employeeUC.On("UpdateMarriageLeave", mockTransaction.EmployeeID, 7).Return(nil)

	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockStatusLeave.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MenstrualApproved() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		MenstrualLeave: 10,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "menstrual",
	}

	// suite.statusLeaveUC.On("FindByNameStatusLeave", mockStatusLeave.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil) // Updated this line

	// Correct the arguments for UpdateAnnualLeave mock
	suite.employeeUC.On("UpdateMenstrualLeave", mockTransaction.EmployeeID, 7).Return(nil)

	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockStatusLeave.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_PaternityApproved() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		PaternityLeave: 10,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "paternity",
	}

	// suite.statusLeaveUC.On("FindByNameStatusLeave", mockStatusLeave.StatusLeaveName).Return(model.StatusLeave{}, nil)
	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil) // Updated this line

	// Correct the arguments for UpdateAnnualLeave mock
	suite.employeeUC.On("PaternityLeave", mockTransaction.EmployeeID, 7).Return(nil)

	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockStatusLeave.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_Rejected() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Rejected",
	}

	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockStatusLeave.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_AnnualReject() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10), // This will exceed available days
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:          "employee_id",
		AnnualLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "annual",
	}

	mockRejectedStatus := model.StatusLeave{
		ID:              "rejected_status_id",
		StatusLeaveName: "Rejected",
	}

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(mockRejectedStatus, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockRejectedStatus.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MaternityReject() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10), // This will exceed available days
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		MaternityLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "maternity",
	}

	mockRejectedStatus := model.StatusLeave{
		ID:              "rejected_status_id",
		StatusLeaveName: "Rejected",
	}

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(mockRejectedStatus, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockRejectedStatus.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MarriageReject() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10), // This will exceed available days
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:            "employee_id",
		MarriageLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "marriage",
	}

	mockRejectedStatus := model.StatusLeave{
		ID:              "rejected_status_id",
		StatusLeaveName: "Rejected",
	}

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(mockRejectedStatus, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockRejectedStatus.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MenstrualReject() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10), // This will exceed available days
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		MenstrualLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "menstrual",
	}

	mockRejectedStatus := model.StatusLeave{
		ID:              "rejected_status_id",
		StatusLeaveName: "Rejected",
	}

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(mockRejectedStatus, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockRejectedStatus.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_PaternityReject() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10), // This will exceed available days
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		PaternityLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "paternity",
	}

	mockRejectedStatus := model.StatusLeave{
		ID:              "rejected_status_id",
		StatusLeaveName: "Rejected",
	}

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)
	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(mockRejectedStatus, nil)
	suite.MockTransactionLeaveRepo.On("UpdateStatus", mockTransaction.ID, mockRejectedStatus.ID).Return(nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.NoError(suite.T(), err)
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_TransactionRetrievalError() {
	mockTransaction := model.TransactionLeave{
		ID: "trx_id",
	}

	expectedError := errors.New("transaction retrieval error")

	// Set up the mock to return an error when GetByIdTxNonDto is called
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(model.TransactionLeave{}, expectedError)

	// Call the method under test
	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	// Check for the expected error
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())

	// Verify that the expected mock method call was made
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_StatusLeaveRetrievalError() {
	mockTransaction := model.TransactionLeave{
		StatusLeaveID: "trx_status_id",
	}

	expectedError := errors.New("transaction retrieval error")

	// Set up the mock to return an error when GetByIdTxNonDto is called
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(model.TransactionLeave{}, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(model.StatusLeave{}, expectedError)

	// Call the method under test
	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	// Check for the expected error
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())

	// Verify that the expected mock method call was made
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_EmployeeRetrievalError() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	expectedError := errors.New("error retrieving employee")

	// Simulate an error when retrieving employee data
	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(model.Employee{}, expectedError)

	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(model.LeaveType{}, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil)

	// UpdateStatus should not be called in case of an error
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_LeaveTypeRetrievalError() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 2),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	expectedError := errors.New("error retrieving employee")

	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(model.LeaveType{}, expectedError)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockStatusLeave.ID).Return(mockStatusLeave, nil)

	// UpdateStatus should not be called in case of an error
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_AnnualReject_RejectedStatusNotFound() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:          "employee_id",
		AnnualLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "annual",
	}

	expectedError := errors.New("rejected status not found")

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(model.StatusLeave{}, expectedError)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MaternityReject_RejectedStatusNotFound() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		MaternityLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "maternity",
	}

	expectedError := errors.New("rejected status not found")

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(model.StatusLeave{}, expectedError)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MarriageReject_RejectedStatusNotFound() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:            "employee_id",
		MarriageLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "marriage",
	}

	expectedError := errors.New("rejected status not found")

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(model.StatusLeave{}, expectedError)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_MenstrualReject_RejectedStatusNotFound() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		MenstrualLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "menstrual",
	}

	expectedError := errors.New("rejected status not found")

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(model.StatusLeave{}, expectedError)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func (suite *TransactionLeaveUseCaseSuite) TestApproveOrRejectLeave_PaternityReject_RejectedStatusNotFound() {
	mockTransaction := model.TransactionLeave{
		ID:            "trx_id",
		EmployeeID:    "employee_id",
		LeaveTypeID:   "leave_type_id",
		StatusLeaveID: "status_leave_id",
		DateStart:     time.Now(),
		DateEnd:       time.Now().Add(time.Hour * 24 * 10),
	}
	mockStatusLeave := model.StatusLeave{
		ID:              "status_leave_id",
		StatusLeaveName: "Approved",
	}

	mockEmployee := model.Employee{
		ID:             "employee_id",
		PaternityLeave: 5,
	}

	mockLeaveType := model.LeaveType{
		ID:            "leave_type_id",
		LeaveTypeName: "paternity",
	}

	expectedError := errors.New("rejected status not found")

	suite.employeeUC.On("FindByIdEmpl", mockTransaction.EmployeeID).Return(mockEmployee, nil)
	suite.leaveTypeUC.On("FindByIdLeaveType", mockTransaction.LeaveTypeID).Return(mockLeaveType, nil)
	suite.statusLeaveUC.On("FindByIdStatusLeave", mockTransaction.StatusLeaveID).Return(mockStatusLeave, nil)
	suite.MockTransactionLeaveRepo.On("GetByIdTxNonDto", mockTransaction.ID).Return(mockTransaction, nil)

	suite.statusLeaveUC.On("FindByNameStatusLeave", "Rejected").Return(model.StatusLeave{}, expectedError)

	err := suite.transactionUC.ApproveOrRejectLeave(mockTransaction)

	assert.EqualError(suite.T(), err, expectedError.Error())
	suite.MockTransactionLeaveRepo.AssertExpectations(suite.T())
	suite.employeeUC.AssertExpectations(suite.T())
	suite.leaveTypeUC.AssertExpectations(suite.T())
	suite.statusLeaveUC.AssertExpectations(suite.T())
}

func TestTransactionLeaveUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TransactionLeaveUseCaseSuite))
}
