package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (uint, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for parsing errors
	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Extract the email from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("failed to extract claims")
	}
	userId, ok := claims["user_id"].(uint)
	if !ok {
		return 0, fmt.Errorf("userId claim not found or invalid")
	}

	// Return the userId and nil error if everything is successful
	return userId, nil
}
