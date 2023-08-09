package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	LeaveTypeUseCase() usecase.LeaveTypeUseCase
	PositionUseCase() usecase.PositionUseCase
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

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
