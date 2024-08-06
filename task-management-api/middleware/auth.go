package middleware

import (
   
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
    UserID string `json:"userID"`
    jwt.StandardClaims
}

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
