package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"
)

type PositionUseCase interface {
	RegisterNewPosition(payload model.Position) error
	FindAllPosition() ([]model.Position, error)
	FindByIdPosition(id string) (model.Position, error)
	UpdatePosition(payload model.Position) error
	DeletePosition(id string) error
	GetByName(name string) (model.Position, error)
}

type positionUseCase struct {
	repo repository.PositionRepository
}

func (p *positionUseCase) RegisterNewPosition(payload model.Position) error {
	if payload.Name == "" {
		return fmt.Errorf("name are required fields")
	}
	isExistManager, _ := p.repo.GetByName(payload.Name)
	if isExistManager.Name == payload.Name {
		return fmt.Errorf("position with name %s exists", payload.Name)
	}
	err := p.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new position: %v", err)
	}
	return nil
}

func (p *positionUseCase) FindAllPosition() ([]model.Position, error) {
	return p.repo.List()
}

func (p *positionUseCase) FindByIdPosition(id string) (model.Position, error) {
	return p.repo.Get(id)
}

func (p *positionUseCase) UpdatePosition(payload model.Position) error {
	fmt.Println("PositionUseCase.UpdatePosition.payload:", payload)
	return p.repo.Update(payload)
}

func (p *positionUseCase) DeletePosition(id string) error {
	return p.repo.Delete(id)
}

func (p *positionUseCase) GetByName(name string) (model.Position, error) {
	return p.repo.GetByName(name)
}

func NewPositionUseCase(repo repository.PositionRepository) PositionUseCase {
	return &positionUseCase{repo: repo}
}
