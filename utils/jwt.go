package utils

import (
	"nanoshell/database/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a new JWT token for a user
func GenerateToken(user models.User) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"admin":    user.Admin,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate signed token
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// DecodeToken validates a JWT token and returns its claims
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate and return claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
