package initializers

import (
	"github.com/joho/godotenv"
)

// LoadEnvVar loads environment variables from a .env file,
// if the file path is not provided, it will load from .env
func LoadEnvVar(envFilePath string) {
	// TODO: Load environment variables
	if envFilePath == "" {
		envFilePath = ".env"
	}
	err := godotenv.Load(envFilePath)
	if err != nil {
		panic("Error loading .env file")
	}
}