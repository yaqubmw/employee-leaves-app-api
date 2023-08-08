package repository

import (
	"database/sql"
	"employeeleave/model"
)

type QuotaLeaveRepository interface {
	BaseRepository[model.QuotaLeave]
}

type quotaLeaveRepository struct {
	db *sql.DB
}


func (q *quotaLeaveRepository) Create(payload model.QuotaLeave) error {
	_, err := q.db.Exec("INSERT INTO quota_leave (id, remaining_quota) VALUES ($1, $2)", payload.ID, payload.RemainingQuota)
	if err != nil {
		return err
	}
	return nil
}


func (q *quotaLeaveRepository) Get(id string) (model.QuotaLeave, error) {
	var quotaLeave model.QuotaLeave
	err := q.db.QueryRow("SELECT id, remaining_quota FROM quota_leave WHERE id=$1", id).Scan(&quotaLeave.ID, &quotaLeave.RemainingQuota)
	if err != nil {
		return model.QuotaLeave{}, err
	}
	return quotaLeave, nil
}


func (q *quotaLeaveRepository) List() ([]model.QuotaLeave, error) {
	rows, err := q.db.Query("SELECT id, remaining_quota FROM quota_leave")
	if err != nil {
		return nil, err
	}

	var quotaLeaves []model.QuotaLeave
	for rows.Next() {
		var quotaLeave model.QuotaLeave
		err := rows.Scan(&quotaLeave.ID, &quotaLeave.RemainingQuota)
		if err != nil {
			return nil, err
		}

		quotaLeaves = append(quotaLeaves, quotaLeave)
	}

	return quotaLeaves, nil
}


func (q *quotaLeaveRepository) Update(payload model.QuotaLeave) error {
	_, err := q.db.Exec("UPDATE quota_leave SET remaining_quota=$1 WHERE id=$2", payload.RemainingQuota, payload.ID)
	if err != nil {
		return err
	}
	return nil
}


func (q *quotaLeaveRepository) Delete(id string) error {
	_, err := q.db.Exec("DELETE FROM quota_leave WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewQuotaLeaveRepository(db *sql.DB) QuotaLeaveRepository {
	return &quotaLeaveRepository{
		db: db,
	}
}
