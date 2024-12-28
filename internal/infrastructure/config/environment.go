package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file if present
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

// CheckEnvVariables checks if all required environment variables are set
func CheckEnvVariables() {
	requiredEnvVars := []string{
		"APP_ENV",
		"DATABASE_NAME",
		"DATABASE_SCHEMA",
		"DATABASE_PORT",
		"DATABASE_HOST",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"RABBITMQ_USER",
		"RABBITMQ_PASSWORD",
		"RABBITMQ_HOST",
		"RABBITMQ_PORT",
		"RABBITMQ_VHOST",
		"VOTE_QUEUE",
		"REDIS_URL",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Environment variable %s is not set", envVar)
		}
	}
}
