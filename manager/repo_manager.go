package manager

import "employeeleave/repository"

type RepoManager interface {
	// semua repo di daftarkan disini

	EmployeeRepo() repository.EmployeeRepository
}

type repoManager struct {
	infra InfraManager
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmplRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
