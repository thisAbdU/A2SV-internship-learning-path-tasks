package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
    jwtKey = []byte("your-secret-key")
)


type Claims struct {
    UserID string `json:"userID"`
    jwt.StandardClaims
}


func GenerateToken(userID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}


func VerifyToken(c *gin.Context) (*Claims, error) {
    tokenString := extractTokenFromHeader(c)
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }
    return claims, nil
}

func Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        _, err := VerifyToken(c)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }
        c.Next()
    }
}


func extractTokenFromHeader(c *gin.Context) string {
    bearerToken := c.GetHeader("Authorization")
    return bearerToken
}
