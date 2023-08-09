package manager

import "employeeleave/repository"

type RepoManager interface {
	// semua repo di daftarkan disini

	EmployeeRepo() repository.EmplRepository
}

type repoManager struct {
	infra InfraManager
}

// UserRepo implements RepoManager.
// func (r *repoManager) UserRepo() repository.UserRepository {
// 	return repository.NewUserRepository(r.infra.Conn())
// }

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmplRepository {
	return repository.NewEmplRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
