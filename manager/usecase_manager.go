package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	EmployeeUseCase() usecase.EmployeeUseCase
	TransactionUseCase() usecase.TransactionLeaveUseCase
	LeaveTypeUseCase() usecase.LeaveTypeUseCase
	StatusLeaveUseCase() usecase.StatusLeaveUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// EmployeeUseCase implements UseCaseManager.
func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

func (u *useCaseManager) TransactionUseCase() usecase.TransactionLeaveUseCase {
	return usecase.NewTransactionLeaveUseCase(u.repoManager.TransactionLeaveRepo(), u.EmployeeUseCase(), u.LeaveTypeUseCase(), u.StatusLeaveUseCase())
}

func (u *useCaseManager) LeaveTypeUseCase() usecase.LeaveTypeUseCase {
	return usecase.NewLeaveTypeUseCase(u.repoManager.LeaveTypeRepo())
}

// StatusLeaveUseCase implements UseCaseManager.
func (r *useCaseManager) StatusLeaveUseCase() usecase.StatusLeaveUseCase {
	return usecase.NewStatusLeaveUseCase(r.repoManager.StatusLeaveRepo())
}
func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
