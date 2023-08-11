package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"employeeleave/utils/common"
	"fmt"
	"time"
)

type TransactionLeaveUseCase interface {
	ApproveOrRejectLeave(payload model.TransactionLeave) error
	ApplyLeave(payload model.TransactionLeave) error
}

type transactionLeaveUseCase struct {
	transactionRepo repository.TransactionRepository
	employeeUC      EmployeeUseCase
	leaveTypeUC     LeaveTypeUseCase
	statusLeaveUC   StatusLeaveUseCase
}

// Persetujuan atau penolakan cuti oleh atasan
// Pengajuan cuti oleh karyawan
func (tl *transactionLeaveUseCase) ApplyLeave(trx model.TransactionLeave) error {

	employee, err := tl.employeeUC.FindByIdEmpl(trx.EmployeeID)
	if err != nil {
		return err
	}

	leaveType, err := tl.leaveTypeUC.FindByIdLeaveType(trx.LeaveTypeID)
	if err != nil {
		return err
	}

	statusLeave, err := tl.statusLeaveUC.FindByNameStatusLeave("Pending")
	if err != nil {
		return err
	}

	historyLeaves := trx.HistoryLeaves
	historyLeaves.Id = common.GenerateID()
	historyLeaves.TransactionLeaveId = trx.ID
	historyLeaves.DateEvent = time.Now()

	trx.EmployeeID = employee.ID
	trx.LeaveTypeID = leaveType.ID
	trx.StatusLeaveID = statusLeave.ID
	trx.SubmissionDate = time.Now()
	trx.HistoryLeaves = historyLeaves

	err = tl.transactionRepo.Create(trx)
	if err != nil {
		return fmt.Errorf("failed to register new transaction %v", err)
	}

	return nil

	// Validasi jumlah cuti yang tersedia
	// if leaveType.QuotaLeave > employee.AvailableLeaveDays {
	// 	return fmt.Errorf("jumlah cuti yang diajukan melebihi sisa cuti yang tersedia")
	// }

	// err = uc.transactionRepo.Create(transaction)
	// if err != nil {
	// 	return err
	// }

	// Kurangi jumlah cuti yang tersedia pada karyawan
	// employee.AvailableLeaveDays -= leaveType.QuotaLeave

	// Update jumlah cuti yang tersedia pada repositori
	// err = uc.employeeRepo.Update(employee)
	// if err != nil {
	// 	return err
	// }

}

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
