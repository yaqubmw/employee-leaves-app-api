package manager

import "employeeleave/repository"

type RepoManager interface {
	StatusLeaveRepo() repository.StatusLeaveRepository
	QuotaLeaveRepo() repository.QuotaLeaveRepository
}

type repoManager struct {
	infra InfraManager
}

// QuotaLeaveRepo implements RepoManager.
func (r *repoManager) QuotaLeaveRepo() repository.QuotaLeaveRepository {
	return repository.NewQuotaLeaveRepository(r.infra.Conn())
}

// StatusLeaveRepo implements RepoManager.
func (r *repoManager) StatusLeaveRepo() repository.StatusLeaveRepository {
	return repository.NewStatusLeaveRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
