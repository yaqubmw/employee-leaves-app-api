package manager

import "employeeleave/repository"

type RepoManager interface {
	// semua repo di daftarkan disini
	LeaveTypeRepo() repository.LeaveTypeRepository
	StatusLeaveRepo() repository.StatusLeaveRepository
	EmployeeRepo() repository.EmployeeRepository
	TransactionLeaveRepo() repository.TransactionRepository
}

type repoManager struct {
	infra InfraManager
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

func (r *repoManager) TransactionLeaveRepo() repository.TransactionRepository {
	return repository.NewTransactionLeaveRepository(r.infra.Conn())
}

func (r *repoManager) LeaveTypeRepo() repository.LeaveTypeRepository {
	return repository.NewLeaveTypeRepository(r.infra.Conn())
}

// StatusLeaveRepo implements RepoManager.
func (r *repoManager) StatusLeaveRepo() repository.StatusLeaveRepository {
	return repository.NewStatusLeaveRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
