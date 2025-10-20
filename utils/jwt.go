package utils

import (
	"time"
    "os"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(getJWTSecret())

func getJWTSecret() string {
    if v := os.Getenv("JWT_SECRET"); v != "" {
        return v
    }
    return "my_secret_key"
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate a token
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
            Issuer:    os.Getenv("JWT_ISSUER"),
            Subject:   "access-token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// Validate token
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
