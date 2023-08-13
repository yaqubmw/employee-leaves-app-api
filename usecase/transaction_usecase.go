package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
	"employeeleave/utils/common"
	"fmt"
	"time"
)

type TransactionLeaveUseCase interface {
	ApplyLeave(payload model.TransactionLeave) error
	FindByEmployeeId(employeeID string) ([]dto.TransactionResponseDto, error)
}

type transactionLeaveUseCase struct {
	transactionRepo repository.TransactionRepository
	employeeUC      EmployeeUseCase
	// positionUC      PositionUseCase
	leaveTypeUC   LeaveTypeUseCase
	statusLeaveUC StatusLeaveUseCase
}

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

func (tl *transactionLeaveUseCase) FindByEmployeeId(employeeID string) ([]dto.TransactionResponseDto, error) {
	return tl.transactionRepo.GetByEmployeeID(employeeID)
}

func NewTransactionLeaveUseCase(transactionRepo repository.TransactionRepository, employeeUC EmployeeUseCase, leaveTypeUC LeaveTypeUseCase, statusLeaveUC StatusLeaveUseCase) TransactionLeaveUseCase {
	return &transactionLeaveUseCase{
		transactionRepo: transactionRepo,
		employeeUC:      employeeUC,
		leaveTypeUC:     leaveTypeUC,
		statusLeaveUC:   statusLeaveUC,
	}
}
