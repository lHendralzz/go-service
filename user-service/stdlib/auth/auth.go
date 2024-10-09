package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Create a struct to represent your claims
type Claims struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, userID int, secretkey string) (string, error) {
	// Create the claims
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string, secretkey string) (Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to verify the token
		return []byte(secretkey), nil
	})

	if err != nil {
		return Claims{}, err
	}

	// Check if the token is valid and if the claims are of type Claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Optionally, you can check additional claims here
		return *claims, nil
	} else {
		return *claims, fmt.Errorf("invalid token")
	}
}
