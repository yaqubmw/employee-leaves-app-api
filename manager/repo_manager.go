package manager

import "employeeleave/repository"

type RepoManager interface {
	// semua repo didaftarkan disini
	RoleRepo() repository.RoleRepository
	HistoryRepo() repository.HistoryRepository
}

type repoManager struct {
	infra InfraManager
}

// RoleRepo implements RepoManager.
func (r *repoManager) RoleRepo() repository.RoleRepository {
	return repository.NewRoleRepository(r.infra.Conn())
}

// HistoryRepo implements RepoManager.
func (r *repoManager) HistoryRepo() repository.HistoryRepository {
	return repository.NewHistoryRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
