package manager

import "employeeleave/repository"

type RepoManager interface {
	LeaveTypeRepo() repository.LeaveTypeRepository
	PositionRepo() repository.PositionRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) PositionRepo() repository.PositionRepository {
	return repository.NewPositionRepository(r.infra.Conn())
}

func (r *repoManager) LeaveTypeRepo() repository.LeaveTypeRepository {
	return repository.NewLeaveTypeRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
