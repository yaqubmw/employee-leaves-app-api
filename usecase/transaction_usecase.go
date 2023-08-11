package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
)

type TransactionLeaveUseCase interface {
	ApproveOrRejectLeave(payload model.TransactionLeave) error
}

type transactionLeaveUseCase struct {
	transactionRepo repository.TransactionRepository
	employeeUC      EmployeeUseCase
	leaveTypeUC     LeaveTypeUseCase
	statusLeaveUC   StatusLeaveUseCase
}

// Persetujuan atau penolakan cuti oleh atasan
func (tl *transactionLeaveUseCase) ApproveOrRejectLeave(trx model.TransactionLeave) error {

	// Get the transaction by ID
	transaction, err := tl.transactionRepo.GetByIdTxNonDto(trx.ID)
	if err != nil {
		return err
	}
	statusLeaveId := transaction.StatusLeaveID
	leaveTypeId := transaction.LeaveTypeID
	employeeId := transaction.EmployeeID

	// If the status is "approved," update the availableLeaveDays for the employee
	if statusLeaveId == "2" {
		leaveType, err := tl.leaveTypeUC.FindByIdLeaveType(leaveTypeId)
		if err != nil {
			return err
		}

		// Calculate the number of leave days

		leaveDays := leaveType.QuotaLeave

		// leaveDays := int(transaction.DateEnd.Sub(transaction.DateStart).Hours()/24) + 1

		// Update the employee's availableLeaveDays
		employee, err := tl.employeeUC.FindByIdEmpl(employeeId)
		if err != nil {
			return err
		}
		employee.AvailableLeaveDays -= leaveDays
		// Update the employee's available leave days in the employee repository
		if err := tl.employeeUC.UpdateAvailableDay(employee.ID, employee.AvailableLeaveDays); err != nil {
			return err
		}
	}

	// Update the status in the transaction repository
	if err := tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID); err != nil {
		return err
	}

	return nil
}

func NewTransactionLeaveUseCase(transactionRepo repository.TransactionRepository, employeeUC EmployeeUseCase, leaveTypeUC LeaveTypeUseCase, statusLeaveUC StatusLeaveUseCase) TransactionLeaveUseCase {
	return &transactionLeaveUseCase{
		transactionRepo: transactionRepo,
		employeeUC:      employeeUC,
		leaveTypeUC:     leaveTypeUC,
		statusLeaveUC:   statusLeaveUC,
	}
}
