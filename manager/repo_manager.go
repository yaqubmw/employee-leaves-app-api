package manager

import (
	"employeeleave/repository"

	"gorm.io/gorm"
)

type RepoManager interface {
	LeaveTypeRepo() repository.LeaveTypeRepository
	PositionRepo() repository.PositionRepository
	StatusLeaveRepo() repository.StatusLeaveRepository
	QuotaLeaveRepo() repository.QuotaLeaveRepository
	RoleRepo() repository.RoleRepository
	HistoryRepo() repository.HistoryRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	db *gorm.DB // Use gorm.DB instead of InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.db)
}

func (r *repoManager) PositionRepo() repository.PositionRepository {
	return repository.NewPositionRepository(r.db)
}

func (r *repoManager) LeaveTypeRepo() repository.LeaveTypeRepository {
	return repository.NewLeaveTypeRepository(r.db)
}

func (r *repoManager) QuotaLeaveRepo() repository.QuotaLeaveRepository {
	return repository.NewQuotaLeaveRepository(r.db)
}

func (r *repoManager) StatusLeaveRepo() repository.StatusLeaveRepository {
	return repository.NewStatusLeaveRepository(r.db)
}

func (r *repoManager) RoleRepo() repository.RoleRepository {
	return repository.NewRoleRepository(r.db)
}

func (r *repoManager) HistoryRepo() repository.HistoryRepository {
	return repository.NewHistoryRepository(r.db)
}

func NewRepoManager(db *gorm.DB) RepoManager {
	return &repoManager{db: db}
}
