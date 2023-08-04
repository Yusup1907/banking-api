package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims interface {
	GenerateToken(email string) (string, error)
	VerifyAccessToken(tokenString string) (string, error)
}

type jwtClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

var KEY = []byte("skajdaslkd")

func GenerateToken(email string) (string, error) {
	now := time.Now().UTC()
	end := now.Add(1 * time.Hour)
	claim := &jwtClaims{
		Email: email,
	}

	claim.IssuedAt = now.Unix()
	claim.ExpiresAt = end.Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := t.SignedString(KEY)
	if err != nil {
		return "", fmt.Errorf("GenerateToken : %w", err)
	}
	return token, nil
}

func VerifyAccessToken(tokenString string) (string, error) {
	claim := &jwtClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return KEY, nil
	})
	if err != nil {
		return "", fmt.Errorf("VerifyAccessToken : %w", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("VerifyAccessToken : Invalid token")
	}
	return claim.Email, nil

}
