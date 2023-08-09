package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	LeaveTypeUseCase() usecase.LeaveTypeUseCase
	PositionUseCase() usecase.PositionUseCase
	StatusLeaveUseCase() usecase.StatusLeaveUseCase
	QuotaLeaveUseCase() usecase.QuotaLeaveUseCase
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.UserUseCase())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func (u *useCaseManager) PositionUseCase() usecase.PositionUseCase {
	return usecase.NewPositionUseCase(u.repoManager.PositionRepo())
}

func (u *useCaseManager) LeaveTypeUseCase() usecase.LeaveTypeUseCase {
	return usecase.NewLeaveTypeUseCase(u.repoManager.LeaveTypeRepo())
}

// QuotaLeaveUseCase implements UseCaseManager.
func (r *useCaseManager) QuotaLeaveUseCase() usecase.QuotaLeaveUseCase {
	return usecase.NewQuotaLeaveUseCase(r.repoManager.QuotaLeaveRepo())
}

// StatusLeaveUseCase implements UseCaseManager.
func (r *useCaseManager) StatusLeaveUseCase() usecase.StatusLeaveUseCase {
	return usecase.NewStatusLeaveUseCase(r.repoManager.StatusLeaveRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
