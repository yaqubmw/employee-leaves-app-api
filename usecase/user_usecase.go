package usecase

import (
	"employeeleave/model"
	"employeeleave/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential) error
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

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
