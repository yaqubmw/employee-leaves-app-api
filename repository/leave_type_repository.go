package repository

import (
	"database/sql"
	"employeeleave/model"
	"fmt"
)

type LeaveTypeRepository interface {
	BaseRepository[model.LeaveType]
	GetByName(name string) (model.LeaveType, error)
}

type leaveTypeRepository struct {
	db *sql.DB
}

func (lt *leaveTypeRepository) Create(payload model.LeaveType) error {
	_, err := lt.db.Exec("INSERT INTO leave_type (id, leave_type_name, quota_leave) VALUES ($1, $2, $3)", payload.ID, payload.LeaveTypeName, payload.QuotaLeave)
	if err != nil {
		return err
	}

	fmt.Println("leave type created sucessfully")
	return nil
}

func (lt *leaveTypeRepository) List() ([]model.LeaveType, error) {
	rows, err := lt.db.Query("SELECT id, leave_type_name, quota_leave FROM leave_type ORDER BY id")
	if err != nil {
		return nil, err
	}

	var leavetypes []model.LeaveType
	for rows.Next() {
		var leave_type model.LeaveType
		err := rows.Scan(&leave_type.ID, &leave_type.LeaveTypeName, &leave_type.QuotaLeave)
		if err != nil {
			return nil, err
		}
		leavetypes = append(leavetypes, leave_type)
	}
	fmt.Println("leave type retrieve all successfully")
	return leavetypes, nil
}

func (lt *leaveTypeRepository) Get(id string) (model.LeaveType, error) {
	var leave_type model.LeaveType
	err := lt.db.QueryRow("SELECT id, leave_type_name, quota_leave FROM leave_type WHERE id=$1", id).Scan(&leave_type.ID, &leave_type.LeaveTypeName, &leave_type.QuotaLeave)
	if err != nil {
		return model.LeaveType{}, err
	}
	return leave_type, nil
}

func (lt *leaveTypeRepository) GetByName(leave_type_name string) (model.LeaveType, error) {
	var leave_type model.LeaveType
	err := lt.db.QueryRow("SELECT id, leave_type_name, quota_leave FROM leave_type WHERE name ILIKE $1", "%"+leave_type_name+"%").Scan(&leave_type.ID, &leave_type.LeaveTypeName, &leave_type.QuotaLeave)
	if err != nil {
		return model.LeaveType{}, err
	}
	return leave_type, nil
}

func (lt *leaveTypeRepository) Update(payload model.LeaveType) error {
	_, err := lt.db.Exec("UPDATE leave_type SET leave_type_name=$1, quota_leave=$2 WHERE id=$3", payload.LeaveTypeName, payload.QuotaLeave, payload.ID)
	if err != nil {
		return err
	}
	fmt.Println("Successfully Updated")
	return nil
}

func (lt *leaveTypeRepository) Delete(id string) error {
	_, err := lt.db.Exec("DELETE FROM leave_type WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewLeaveTypeRepository(db *sql.DB) LeaveTypeRepository {
	return &leaveTypeRepository{db: db}
}
