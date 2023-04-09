package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Define a expiração do token

	authKey := os.Getenv("AUTH_KEY")

	tokenString, err := token.SignedString([]byte(authKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
