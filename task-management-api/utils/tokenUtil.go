package utils

import (
	"task-management-api/config"
	"task-management-api/domain/entities"
	"time"

	"github.com/golang-jwt/jwt"
)

type Utils interface {
	GenerateToken(userID string) (string, error)
}

type TokenUtil struct {
	environment *config.Environment
}

func (t *TokenUtil) GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &entities.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	var jwtKey = config.GetJwtKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewTokenUtil(environment *config.Environment) *TokenUtil {
	return &TokenUtil{
		environment: environment,
	}
}

func GetJwtKey() []byte {
	return config.GetJwtKey()
}
