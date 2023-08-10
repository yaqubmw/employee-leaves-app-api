package manager

import "employeeleave/repository"

type RepoManager interface {
	LeaveTypeRepo() repository.LeaveTypeRepository
	PositionRepo() repository.PositionRepository
	StatusLeaveRepo() repository.StatusLeaveRepository
	QuotaLeaveRepo() repository.QuotaLeaveRepository
	RoleRepo() repository.RoleRepository
	// HistoryRepo() repository.HistoryRepository
	EmployeeRepo() repository.EmployeeRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmplRepository(r.infra.Conn())
}

func (r *repoManager) PositionRepo() repository.PositionRepository {
	return repository.NewPositionRepository(r.infra.Conn())
}

func (r *repoManager) LeaveTypeRepo() repository.LeaveTypeRepository {
	return repository.NewLeaveTypeRepository(r.infra.Conn())
}

// QuotaLeaveRepo implements RepoManager.
func (r *repoManager) QuotaLeaveRepo() repository.QuotaLeaveRepository {
	return repository.NewQuotaLeaveRepository(r.infra.Conn())
}

// StatusLeaveRepo implements RepoManager.
func (r *repoManager) StatusLeaveRepo() repository.StatusLeaveRepository {
	return repository.NewStatusLeaveRepository(r.infra.Conn())
}

// RoleRepo implements RepoManager.
func (r *repoManager) RoleRepo() repository.RoleRepository {
	return repository.NewRoleRepository(r.infra.Conn())
}

// // HistoryRepo implements RepoManager.
// func (r *repoManager) HistoryRepo() repository.HistoryRepository {
// 	return repository.NewHistoryRepository(r.infra.Conn())
// }

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
