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
    // Set expiration time for the token (24 hours in this example)
    expirationTime := time.Now().Add(24 * time.Hour)

    // Create custom claims
    claims := &Claims{
        UserID: s,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    // Create a token with the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with a secret key
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
