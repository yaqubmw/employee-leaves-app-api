package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	EmployeeUseCase() usecase.EmployeeUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// EmployeeUseCase implements UseCaseManager.
func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
