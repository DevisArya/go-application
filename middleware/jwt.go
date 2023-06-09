package middleware

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CreateToken(userId int, name string) (string, error) {
	payload := jwt.MapClaims{}
	payload["userId"] = userId
	payload["name"] = name
	payload["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")

	return token.SignedString([]byte(secretKey))
}
