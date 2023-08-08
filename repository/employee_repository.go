package repository

import (
	"database/sql"
	"employeeleave/model"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	// BaseRepositoryPaging[model.Employee]
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) Create(payload model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee (id, position_id. manager_id name, phone_number, email, address) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		payload.ID, payload.PositionID, payload.ManagerID, payload.Name, payload, payload.PhoneNumber, payload.Email, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) Delete(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) Get(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.QueryRow("SELECT id, position_id. manager_id name, phone_number, email, address FROM employee WHERE id=$1", id).
		Scan(&employee.ID, &employee.PositionID, &employee.ManagerID, &employee.Name, &employee.PhoneNumber, &employee.Email,
			&employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil

}

func (e *employeeRepository) List() ([]model.Employee, error) {
	rows, err := e.db.Query("SELECT id, position_id. manager_id name, phone_number, email, address FROM employee")
	if err != nil {
		return nil, err
	}
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.ID, &employee.PositionID, &employee.ManagerID, &employee.Name, &employee.PhoneNumber, &employee.Email,
			&employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *employeeRepository) Update(payload model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET position_id = $2, manager_id = $3, name = $4, phone_number = $5, email = $6, address = $7 WHERE id = $1",
		payload.ID, payload.PositionID, payload.ManagerID, payload.Name, payload.PhoneNumber, payload.Email, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db}
}
