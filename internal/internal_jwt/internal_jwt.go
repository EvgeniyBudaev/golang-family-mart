package internal_jwt

import (
	"github.com/EvgeniyBudaev/golang-family-mart/internal/models"
	"github.com/golang-jwt/jwt"
	"os"
)

func getSecretKey() string {
	return os.Getenv("SECRET")
}

func GenerateJWT(user models.User) (string, error) {
	mySigningKey := []byte(getSecretKey())

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}
