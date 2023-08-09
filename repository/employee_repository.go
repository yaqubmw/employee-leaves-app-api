package repository

import (
	"database/sql"
	"employeeleave/model"
)

// Karena dia sebuah interface, maka wajib kita implementasikan semuanya
// mulai dari Create s.d Delete
type EmplRepository interface {
	BaseRepository[model.Employee]
	GetByName(name string) (model.Employee, error)
}

type emplRepository struct {
	db *sql.DB
}

// Method -> ada sebuah receiver ((u *uomRepository))
func (e *emplRepository) Create(payload model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee (id, name, phonenumber, email, address) VALUES ($1, $2, $3, $4, $5)", payload.ID, payload.Name, payload.PhoneNumber, payload.Email, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (e *emplRepository) List() ([]model.Employee, error) {
	rows, err := e.db.Query("SELECT id, name, phonenumber, email, address FROM employee")
	if err != nil {
		return nil, err
	}

	var empls []model.Employee
	for rows.Next() {
		var empl model.Employee
		err := rows.Scan(&empl.ID, &empl.Name, &empl.PhoneNumber, &empl.Email, &empl.Address)
		if err != nil {
			return nil, err
		}
		empls = append(empls, empl)
	}
	return empls, nil
}

func (e *emplRepository) Get(id string) (model.Employee, error) {
	var empl model.Employee
	err := e.db.QueryRow("SELECT id, name, phonenumber, email, address FROM employee WHERE id=$1", id).Scan(&empl.ID, &empl.Name, &empl.PhoneNumber, &empl.Email, &empl.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return empl, nil
}

func (e *emplRepository) GetByName(name string) (model.Employee, error) {
	var empl model.Employee
	// LIKE => case sensitive e.g L l (ngaruh)
	// ILIKE => in case sensitibe e.g L l (tidak ngaruh) (hanya ada di postgre)
	err := e.db.QueryRow("SELECT id, name, phonenumber, email, address FROM employee WHERE name ILIKE $1", "%"+name+"%").Scan(&empl.ID, &empl.Name, &empl.PhoneNumber, &empl.Email, &empl.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return empl, nil
}

func (e *emplRepository) Update(payload model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET name=$1, phonenumber=$3, email=$4, address=$5 WHERE id=$2", payload.Name, payload.ID, payload.PhoneNumber, payload.Email, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (e *emplRepository) Delete(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Constructor
func NewEmplRepository(db *sql.DB) *emplRepository {
	return &emplRepository{db: db}
}
