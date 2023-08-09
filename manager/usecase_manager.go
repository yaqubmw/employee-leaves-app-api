package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	RoleUseCase() usecase.RoleUseCase
	HistoryUseCase() usecase.HistoryUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// RoleUseCase implements UseCaseManager.
func (u *useCaseManager) RoleUseCase() usecase.RoleUseCase {
	return usecase.NewRoleUseCase(u.repoManager.RoleRepo())
}

// HistoryUseCase implements UseCaseManager.
func (u *useCaseManager) HistoryUseCase() usecase.HistoryUseCase {
	return usecase.NewHistoryUseCase(u.repoManager.HistoryRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
