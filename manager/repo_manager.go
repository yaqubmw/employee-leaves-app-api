package manager

import "employeeleave/repository"

type RepoManager interface {
	LeaveTypeRepo() repository.LeaveTypeRepository
	PositionRepo() repository.PositionRepository
	StatusLeaveRepo() repository.StatusLeaveRepository
	RoleRepo() repository.RoleRepository
	HistoryRepo() repository.HistoryRepository
	EmployeeRepo() repository.EmployeeRepository
	UserRepo() repository.UserRepository
	TransactionLeaveRepo() repository.TransactionRepository
}

type repoManager struct {
	infra InfraManager
}

// TransactionLeaveRepo implements RepoManager.
func (r *repoManager) TransactionLeaveRepo() repository.TransactionRepository {
	return repository.NewTransactionLeaveRepository(r.infra.Conn())
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

// StatusLeaveRepo implements RepoManager.
func (r *repoManager) StatusLeaveRepo() repository.StatusLeaveRepository {
	return repository.NewStatusLeaveRepository(r.infra.Conn())
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
