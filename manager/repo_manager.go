package manager

import "employeeleave/repository"

type RepoManager interface {
	// semua repo di daftarkan disini

	EmployeeRepo() repository.EmployeeRepository
}

type repoManager struct {
	infra InfraManager
}

// UserRepo implements RepoManager.
// func (r *repoManager) UserRepo() repository.UserRepository {
// 	return repository.NewUserRepository(r.infra.Conn())
// }

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
