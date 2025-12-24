package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var AccessSecret = []byte("ACCESS_SECRET_CHANGE_ME")
var RefreshSecret = []byte("REFRESH_SECRET_CHANGE_ME")

// ==========================
// Generate Access Token (15 min)
// ==========================
func GenerateAccessToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessSecret)
}

// ==========================
// Generate Refresh Token (7 days)
// ==========================
func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshSecret)
}

// ==========================
// Parse Access Token
// ==========================
func ParseAccessToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return AccessSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, token.Claims.(jwt.MapClaims), nil
}

// ==========================
// Parse Refresh Token
// ==========================
func ParseRefreshToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return RefreshSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, token.Claims.(jwt.MapClaims), nil
}
