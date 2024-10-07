package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Create a struct to represent your claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, secretkey string) (string, error) {
	// Create the claims
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
