package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	EmployeeUseCase() usecase.EmplUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// AuthUseCase implements UseCaseManager.

// EmployeeUseCase implements UseCaseManager.
func (u *useCaseManager) EmployeeUseCase() usecase.EmplUseCase {
	return usecase.NewEmplUseCase(u.repoManager.EmployeeRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
