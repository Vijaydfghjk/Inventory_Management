package token_stuff

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("YourSecretKey")

type JWTClaim struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT generates token
func GenerateJWT(userID int, email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT parses token and returns claims
func ValidateJWT(tokenStr string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, err
	}
	return claims, nil
}
