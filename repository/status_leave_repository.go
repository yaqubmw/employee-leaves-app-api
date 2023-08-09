package repository

import (
	"database/sql"
	"employeeleave/model"
)

type StatusLeaveRepository interface {
	BaseRepository[model.StatusLeave]
	GetByName(name string) (model.StatusLeave, error)
}

type statusLeaveRepository struct {
	db *sql.DB
}

func (s *statusLeaveRepository) Create(payload model.StatusLeave) error {
	_, err := s.db.Exec("INSERT INTO status_leave (id, status_leave_name) VALUES ($1, $2)", payload.ID, payload.StatusLeaveName)
	if err != nil {
		return err
	}
	return nil
}

func (s *statusLeaveRepository) Get(id string) (model.StatusLeave, error) {
	var statusLeave model.StatusLeave
	err := s.db.QueryRow("SELECT id, status_leave_name FROM status_leave WHERE id=$1", id).Scan(&statusLeave.ID, &statusLeave.StatusLeaveName)
	if err != nil {
		return model.StatusLeave{}, err
	}
	return statusLeave, nil
}

func (s *statusLeaveRepository) GetByName(status string) (model.StatusLeave, error) {
	var statusLeave model.StatusLeave
	err := s.db.QueryRow("SELECT id, status_leave_name FROM status_leave WHERE status_leave_name ILIKE $1", "%"+status+"%").Scan(&statusLeave.ID, &statusLeave.StatusLeaveName)
	if err != nil {
		return model.StatusLeave{}, err
	}
	return statusLeave, nil
}

func (s *statusLeaveRepository) List() ([]model.StatusLeave, error) {
	rows, err := s.db.Query("SELECT id, status_leave_name FROM status_leave")
	if err != nil {
		return nil, err
	}

	var statusLeaves []model.StatusLeave
	for rows.Next() {
		var statusLeave model.StatusLeave
		err := rows.Scan(&statusLeave.ID, &statusLeave.StatusLeaveName)
		if err != nil {
			return nil, err
		}

		statusLeaves = append(statusLeaves, statusLeave)
	}

	return statusLeaves, nil
}

func (s *statusLeaveRepository) Update(payload model.StatusLeave) error {
	_, err := s.db.Exec("UPDATE status_leave SET status_leave_name=$1 WHERE id=$2", payload.StatusLeaveName, payload.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *statusLeaveRepository) Delete(id string) error {
	_, err := s.db.Exec("DELETE FROM status_leave WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewStatusLeaveRepository(db *sql.DB) StatusLeaveRepository {
	return &statusLeaveRepository{
		db: db,
	}
}
