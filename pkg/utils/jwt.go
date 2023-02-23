package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sayuen0/go-to-gym/config"
)

const (
	jwtExpiresMinutes = 3 * 60
)

// Claims is a wrapper of jwt.StandardClaims with auth info
type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

// GenerateJWTToken returns a Claims object
func GenerateJWTToken(email string, userID string, config *config.Config) (string, error) {
	claims := &Claims{
		Email: email,
		ID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * jwtExpiresMinutes).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Server.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
