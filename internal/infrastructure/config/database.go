package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"bbb-voting-service/internal/infrastructure/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitDB initializes the database connection and runs migrations.
func InitDB() *gorm.DB {
	databaseUser := getEnv("DATABASE_USER")
	databasePassword := getEnv("DATABASE_PASSWORD")
	databaseHost := getEnv("DATABASE_HOST")
	databasePort := getEnv("DATABASE_PORT")
	databaseName := getEnv("DATABASE_NAME")
	databaseSchema := getEnv("DATABASE_SCHEMA")
	appEnv := getEnv("APP_ENV")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s", databaseUser, databasePassword, databaseHost, databasePort)

	if appEnv == "development" {
		createDatabase(databaseURL, databaseName)
		createSchema(databaseURL, databaseName, databaseSchema)
	}

	// Construct the database URL with the schema after creating the database and schema
	databaseURLWithSchema := fmt.Sprintf("%s/%s?search_path=%s&sslmode=disable", databaseURL, databaseName, databaseSchema)
	db, err := gorm.Open(postgres.Open(databaseURLWithSchema), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s.", databaseSchema),
		},
	})
	checkError("Failed to connect to the database", err)

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	checkError("Failed to create uuid-ossp extension", err)

	err = db.AutoMigrate(&models.VoteModel{}, &models.ParticipantModel{})
	checkError("Failed to migrate the database schema", err)

	return db
}

// createDatabase creates the database if it doesn't exist.
func createDatabase(databaseURL, databaseName string) {
	// Ensure the URL uses the default 'postgres' database
	databaseURLWithDefaultDB := fmt.Sprintf("%s/postgres?sslmode=disable", databaseURL)
	db, err := sql.Open("postgres", databaseURLWithDefaultDB)
	checkError("Failed to connect to the database server", err)
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil && err.Error() != fmt.Sprintf("pq: database \"%s\" already exists", databaseName) {
		checkError("Failed to create database", err)
	}
}

// createSchema creates the schema if it doesn't exist.
func createSchema(databaseURL, databaseName, schemaName string) {
	databaseURLWithDB := fmt.Sprintf("%s/%s?sslmode=disable", databaseURL, databaseName)
	db, err := sql.Open("postgres", databaseURLWithDB)
	checkError("Failed to connect to the database", err)
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
	checkError("Failed to create schema", err)
}

func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("%s is not set", key)
	}

	return value
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
