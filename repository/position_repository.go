package repository

import (
	"database/sql"
	"employeeleave/model"
	"fmt"
)

type PositionRepository interface {
	BaseRepository[model.Position]
	GetByName(name string) (model.Position, error)
}

type positionRepository struct {
	db *sql.DB
}

func (p *positionRepository) Create(payload model.Position) error {
	_, err := p.db.Exec("INSERT INTO position (id, name, is_manager) VALUES ($1, $2, $3)", payload.ID, payload.Name, payload.IsManager)
	if err != nil {
		return err
	}

	fmt.Println("position created sucessfully")
	return nil
}

func (p *positionRepository) List() ([]model.Position, error) {
	rows, err := p.db.Query("SELECT id, name, is_manager FROM position ORDER BY id")
	if err != nil {
		return nil, err
	}

	var positions []model.Position
	for rows.Next() {
		var position model.Position
		err := rows.Scan(&position.ID, &position.Name, &position.IsManager)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	fmt.Println("position retrieve all successfully")
	return positions, nil
}

func (p *positionRepository) Get(id string) (model.Position, error) {
	var position model.Position
	err := p.db.QueryRow("SELECT id, name, is_manager FROM position WHERE id=$1", id).Scan(&position.ID, &position.Name, &position.IsManager)
	if err != nil {
		return model.Position{}, err
	}
	return position, nil
}

func (p *positionRepository) GetByName(name string) (model.Position, error) {
	var position model.Position
	err := p.db.QueryRow("SELECT id, name, is_manager FROM position WHERE name ILIKE $1", "%"+name+"%").Scan(&position.ID, &position.Name, &position.IsManager)
	if err != nil {
		return model.Position{}, err
	}
	return position, nil
}

func (p *positionRepository) Update(payload model.Position) error {
	_, err := p.db.Exec("UPDATE position SET name=$1, is_manager=$2 WHERE id=$3", payload.Name, payload.IsManager, payload.ID)
	if err != nil {
		return err
	}
	fmt.Println("Successfully Updated")
	return nil
}

func (p *positionRepository) Delete(id string) error {
	_, err := p.db.Exec("DELETE FROM position WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewPositionRepository(db *sql.DB) PositionRepository {
	return &positionRepository{db: db}
}
