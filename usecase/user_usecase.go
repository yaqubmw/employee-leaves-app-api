package usecase

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential) error
	FindAllUser(requesPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error)
	FindByUsername(username string) (model.UserCredential, error)
	FindByUsernamePassword(username, password string) (model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload model.UserCredential) error {
	// buat hash
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser(requesPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error) {
	return u.repo.Paging(dto.PaginationParam{})
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUsername(username string) (model.UserCredential, error) {
	return u.repo.GetByUsername(username)
}

// FindByUsernamePassword implements UserUseCase.
func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.UserCredential, error) {
	return u.repo.GetByUsernamePassword(username, password)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
