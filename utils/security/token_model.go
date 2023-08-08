package security

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenMyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string
	Services []string
}
