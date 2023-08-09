package manager

import "employeeleave/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// AuthUseCase implements UseCaseManager.
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.UserUseCase())
}

// UserUseCase implements UseCaseManager.
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
