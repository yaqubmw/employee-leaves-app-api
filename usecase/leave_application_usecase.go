package usecase

// import (
// 	"employeeleave/model"
// 	"employeeleave/model/dto"
// 	"employeeleave/repository"
// 	"fmt"
// 	"time"
// )

// type LeaveApplicationUseCase interface {
// 	ApplyLeave(employeeID, leaveTypeID string, dateStart, dateEnd time.Time, typeOfDay, reason string) error
// 	ApproveLeave(transactionID, managerID string) error
// 	ApproveLeaveByHC(transactionID string) error
// 	GetLeaveStatus(employeeID string) ([]dto.TransactionResponseDto, error)
// }

// type leaveApplicationUseCase struct {
// 	transactionRepo repository.TransactionRepository
// 	employeeRepo    repository.EmployeeRepository
// 	positionRepo    repository.PositionRepository
// 	leaveTypeRepo   repository.LeaveTypeRepository
// 	statusLeaveRepo repository.StatusLeaveRepository
// }

// // Pengajuan cuti oleh karyawan
// func (uc *leaveApplicationUseCase) ApplyLeave(employeeID, leaveTypeID string, dateStart, dateEnd time.Time, typeOfDay, reason string) error {
// 	employee, err := uc.employeeRepo.Get(employeeID)
// 	if err != nil {
// 		return err
// 	}

// 	leaveType, err := uc.leaveTypeRepo.Get(leaveTypeID)
// 	if err != nil {
// 		return err
// 	}

// 	statusLeave, err := uc.statusLeaveRepo.GetByName("Pending")
// 	if err != nil {
// 		return err
// 	}

// 	// Validasi jumlah cuti yang tersedia
// 	if leaveType.QuotaLeave > employee.AvailableLeaveDays {
// 		return fmt.Errorf("jumlah cuti yang diajukan melebihi sisa cuti yang tersedia")
// 	}

// 	transaction := model.TransactionLeave{
// 		EmployeeID:     employeeID,
// 		LeaveTypeID:    leaveTypeID,
// 		StatusLeaveID:  statusLeave.ID,
// 		DateStart:      dateStart,
// 		DateEnd:        dateEnd,
// 		TypeOfDay:      typeOfDay,
// 		Reason:         reason,
// 		SubmissionDate: time.Now(),
// 	}

// 	err = uc.transactionRepo.Create(transaction)
// 	if err != nil {
// 		return err
// 	}

// 	// Kurangi jumlah cuti yang tersedia pada karyawan
// 	employee.AvailableLeaveDays -= leaveType.QuotaLeave

// 	// Update jumlah cuti yang tersedia pada repositori
// 	err = uc.employeeRepo.Update(employee)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Persetujuan atau penolakan cuti oleh atasan
// func (uc *leaveApplicationUseCase) ApproveOrRejectLeave(transactionID, managerID string, isApproved bool) error {
// 	transaction, err := uc.transactionRepo.GetByID(transactionID)
// 	if err != nil {
// 		return err
// 	}

// 	manager, err := uc.positionRepo.Get(managerID)
// 	if err != nil {
// 		return err
// 	}

// 	if !manager.IsManager {
// 		return fmt.Errorf("pegawai yang menyetujui/mentolak bukan atasan")
// 	}

// 	var statusName string
// 	if isApproved {
// 		statusName = "Approved"
// 	} else {
// 		statusName = "Rejected"
// 	}

// 	statusLeave, err := uc.statusLeaveRepo.GetByName(statusName)
// 	if err != nil {
// 		return err
// 	}

// 	transaction.StatusLeave.ID = statusLeave.ID

// 	err = uc.transactionRepo.UpdateStatus(transactionID, transaction.StatusLeave.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // update cuti oleh HC
// func (uc *leaveApplicationUseCase) ApproveLeaveByHC(transactionID string) error {
// 	transaction, err := uc.transactionRepo.GetByID(transactionID)
// 	if err != nil {
// 		return err
// 	}

// 	statusLeave, err := uc.statusLeaveRepo.GetByName("Approved")
// 	if err != nil {
// 		return err
// 	}

// 	transaction.StatusLeave.ID = statusLeave.ID

// 	err = uc.transactionRepo.UpdateStatus(transactionID, transaction.StatusLeave.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Mendapatkan informasi status cuti untuk pegawai
// func (uc *leaveApplicationUseCase) GetLeaveStatusForEmployee(employeeID string) ([]dto.TransactionResponseDto, error) {
// 	transactions, err := uc.transactionRepo.GetByEmployeeID(employeeID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transactions, nil
// }

// // Menampilkan informasi status cuti kepada pegawai
// func DisplayLeaveStatusToEmployee(transactions []dto.TransactionResponseDto) {
// 	for _, transaction := range transactions {
// 		fmt.Printf("ID: %s, Tanggal Mulai: %s, Status: %s\n", transaction.ID, transaction.DateStart, transaction.StatusLeave.StatusLeaveName)
// 	}
// }

// // Pengiriman pemberitahuan kepada pegawai tentang status cuti
// func NotifyEmployeeAboutLeaveStatus(employeeID string, uc *leaveApplicationUseCase) {
// 	transactions, err := uc.GetLeaveStatusForEmployee(employeeID)
// 	if err != nil {
// 		return
// 	}
// 	DisplayLeaveStatusToEmployee(transactions)
// }
