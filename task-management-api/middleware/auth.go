package middleware

import (
	"task-management-api/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
    UserID string `json:"userID"`
    jwt.StandardClaims
}

var jwtKey = config.GetJwtKey();

func GenerateToken(s string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        UserID: s,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
