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

func InitDB() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	databaseSchema := os.Getenv("DATABASE_SCHEMA")
	if databaseSchema == "" {
		log.Fatal("DATABASE_SCHEMA is not set")
	}

	// Create the schema if it doesn't exist
	createSchema(databaseURL, databaseSchema)

	// Append the schema to the database URL
	databaseURLWithSchema := fmt.Sprintf("%s?search_path=%s&sslmode=disable", databaseURL, databaseSchema)

	database, err := gorm.Open(postgres.Open(databaseURLWithSchema), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s.", databaseSchema), // Set the schema
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create the uuid-ossp extension if it doesn't exist
	err = database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err)
	}

	// Run database migrations using gorm
	err = database.AutoMigrate(&models.VoteModel{}, &models.ParticipantModel{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	return database
}

func createSchema(databaseURL, schemaName string) {
	databaseURLWithSSLMode := fmt.Sprintf("%s?sslmode=disable", databaseURL)

	// Open a connection to the database
	db, err := sql.Open("postgres", databaseURLWithSSLMode)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create the schema if it doesn't exist
	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
}
