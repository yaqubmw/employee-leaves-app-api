package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	StatusLeaveUseCase() usecase.StatusLeaveUseCase
	QuotaLeaveUseCase() usecase.QuotaLeaveUseCase
}

type useCaseManager struct {
	repoManager RepoManager
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
	return &useCaseManager{
		repoManager: repoManager,
	}
}
