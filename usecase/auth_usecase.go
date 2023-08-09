package usecase

import (
	"employeeleave/utils/security"
	"fmt"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	usecase UserUseCase
}

func (a *authUseCase) Login(username, password string) (string, error) {
	user, err := a.usecase.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username and password")
	}

	// mekanisme jika user itu ada akan membalikan sebuah token
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}

func NewAuthUseCase(usecase UserUseCase) AuthUseCase {
	return &authUseCase{usecase: usecase}
}
