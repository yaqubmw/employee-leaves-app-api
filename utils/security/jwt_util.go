package security

import (
	"employeeleave/config"
	"employeeleave/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user model.UserCredential) (string, error) {
	cfg, _ := config.NewConfig()
	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
		// role
		// services
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %v", err)
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, _ := config.NewConfig()
	// digunakan untuk mem-Parse token yang dikirimkan dari CLIENT
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// pengecekan sebuah method yang digunakan
		// validasi tanda tangan seperti (SIGNINING METHOD) yang digunakan yaitu HS256
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signing method")
		}
		// kita kembalikan validasi dari konfigurasi yang sudah di validasi di atas
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	// cek claims yang sudah didaftarkan sebelumnya
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
