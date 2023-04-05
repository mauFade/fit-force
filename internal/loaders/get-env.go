package loaders

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
