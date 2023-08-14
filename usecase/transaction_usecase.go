package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
	"employeeleave/utils/common"
	"employeeleave/utils/helper"
	"fmt"
	"time"
)

type TransactionLeaveUseCase interface {
	ApplyLeave(payload model.TransactionLeave) error
	FindById(id string) (model.TransactionLeave, error)
	FindByIdEmpl(id string) ([]model.TransactionLeave, error)
	FindAllEmpl(requesPaging dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error)
	ApproveOrRejectLeave(payload model.TransactionLeave) error
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

}

func (t *transactionLeaveUseCase) FindAllEmpl(requesPaging dto.PaginationParam) ([]dto.TransactionResponseDto, dto.Paging, error) {
	return t.transactionRepo.Paging(requesPaging)
}

func (t *transactionLeaveUseCase) FindById(id string) (model.TransactionLeave, error) {
	return t.transactionRepo.GetByID(id)
}

func (t *transactionLeaveUseCase) FindByIdEmpl(id string) ([]model.TransactionLeave, error) {
	return t.transactionRepo.GetByEmployeeID(id)
}

func (tl *transactionLeaveUseCase) ApproveOrRejectLeave(trx model.TransactionLeave) error {

	// Get the transaction by ID
	transaction, err := tl.transactionRepo.GetByIdTxNonDto(trx.ID)
	if err != nil {
		return err
	}
	statusLeaveTx, err := tl.statusLeaveUC.FindByIdStatusLeave(trx.StatusLeaveID)
	if err != nil {
		return err
	}
	statusLeaveName := statusLeaveTx.StatusLeaveName

	leaveTypeId := transaction.LeaveTypeID
	employeeId := transaction.EmployeeID
	leaveDays := int(transaction.DateEnd.Sub(transaction.DateStart).Hours()/24) + 1

	// If the status is "approved," update the availableLeaveDays for the employee
	if statusLeaveName == "Approved" {
		leaveType, err := tl.leaveTypeUC.FindByIdLeaveType(leaveTypeId)
		if err != nil {
			return err
		}
		leaveTypeName := leaveType.LeaveTypeName

		employee, err := tl.employeeUC.FindByIdEmpl(employeeId)
		if err != nil {
			return err
		}
		if helper.MatchKeyword(leaveTypeName) == "annual" {

			// auto reject bila sisa cuti lebih kecil dari pengajuan cuti
			if employee.AnnualLeave < leaveDays {
				reject, err := tl.statusLeaveUC.FindByNameStatusLeave("Rejected")
				if err != nil {
					return err
				}
				tl.transactionRepo.UpdateStatus(trx.ID, reject.ID)

			} else {
				// mengurangi sisa cuti
				employee.AnnualLeave -= leaveDays
				tl.employeeUC.UpdateAnnualLeave(employee.ID, employee.AnnualLeave)
				tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID)

			}

		}
		if helper.MatchKeyword(leaveTypeName) == "maternity" {

			// auto reject bila sisa cuti lebih kecil dari pengajuan cuti
			if employee.MaternityLeave < leaveDays {
				reject, err := tl.statusLeaveUC.FindByNameStatusLeave("Rejected")
				if err != nil {
					return err
				}
				tl.transactionRepo.UpdateStatus(trx.ID, reject.ID)

			} else {
				// mengurangi sisa cuti
				employee.MaternityLeave -= leaveDays
				tl.employeeUC.UpdateMaternityLeave(employee.ID, employee.MaternityLeave)
				tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID)

			}
		}
		if helper.MatchKeyword(leaveTypeName) == "marriage" {

			// auto reject bila sisa cuti lebih kecil dari pengajuan cuti
			if employee.MarriageLeave < leaveDays {
				reject, err := tl.statusLeaveUC.FindByNameStatusLeave("Rejected")
				if err != nil {
					return err
				}
				tl.transactionRepo.UpdateStatus(trx.ID, reject.ID)

			} else {
				// mengurangi sisa cuti
				employee.MarriageLeave -= leaveDays
				tl.employeeUC.UpdateMarriageLeave(employee.ID, employee.MarriageLeave)
				tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID)

			}
		}
		if helper.MatchKeyword(leaveTypeName) == "menstrual" {

			// auto reject bila sisa cuti lebih kecil dari pengajuan cuti
			if employee.MenstrualLeave < leaveDays {
				reject, err := tl.statusLeaveUC.FindByNameStatusLeave("Rejected")
				if err != nil {
					return err
				}
				tl.transactionRepo.UpdateStatus(trx.ID, reject.ID)

			} else {
				// mengurangi sisa cuti
				employee.MenstrualLeave -= leaveDays
				tl.employeeUC.UpdateMenstrualLeave(employee.ID, employee.MenstrualLeave)
				tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID)

			}
		}

		if helper.MatchKeyword(leaveTypeName) == "paternity" {

			// auto reject bila sisa cuti lebih kecil dari pengajuan cuti
			if employee.PaternityLeave < leaveDays {
				reject, err := tl.statusLeaveUC.FindByNameStatusLeave("Rejected")
				if err != nil {
					return err
				}
				tl.transactionRepo.UpdateStatus(trx.ID, reject.ID)

			} else {
				// mengurangi sisa cuti
				employee.PaternityLeave -= leaveDays
				tl.employeeUC.PaternityLeave(employee.ID, employee.PaternityLeave)
				tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID)

			}
		}
	} else {
		tl.transactionRepo.UpdateStatus(trx.ID, trx.StatusLeaveID)
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
