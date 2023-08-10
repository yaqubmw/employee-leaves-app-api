package repository

import (
	"employeeleave/model"
	"fmt"

	"gorm.io/gorm"
)

type PositionRepository interface {
	BaseRepository[model.Position]
	GetByName(name string) (model.Position, error)
}

type positionRepository struct {
	db *gorm.DB
}

func (p *positionRepository) Create(payload model.Position) error {
	result := p.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("position created successfully")
	return nil
}

func (p *positionRepository) List() ([]model.Position, error) {
	var positions []model.Position
	result := p.db.Find(&positions)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("position retrieve all successfully")
	return positions, nil
}

func (p *positionRepository) Get(id string) (model.Position, error) {
	var position model.Position
	result := p.db.First(&position, id)
	if result.Error != nil {
		return model.Position{}, result.Error
	}
	return position, nil
}

func (p *positionRepository) GetByName(name string) (model.Position, error) {
	var position model.Position
	result := p.db.Where("name ILIKE ?", "%"+name+"%").First(&position)
	if result.Error != nil {
		return model.Position{}, result.Error
	}
	return position, nil
}

func (p *positionRepository) Update(payload model.Position) error {
	result := p.db.Save(&payload)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Successfully Updated")
	return nil
}

func (p *positionRepository) Delete(id string) error {
	result := p.db.Delete(&model.Position{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return &positionRepository{db: db}
}
