package main

import (
	"log"
	"os"
	"strings"
	"time"

	_ "bbb-voting-service/docs" // Import generated Swagger docs
	"bbb-voting-service/internal/infrastructure"
	"bbb-voting-service/internal/infrastructure/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title BBB Voting System API
// @version 1.0
// @description This is the API for the BBB Voting System.
// @host localhost:8080
// @BasePath /
// @schemes http https

// @Summary Health check
// @Description Returns the service's status
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health [get]
func main() {
	// Load environment variables
	config.LoadEnv()

	// Check required environment variables
	config.CheckEnvVariables()

	// Get the port from the environment variable
	port := getPort()

	// Get the allowed origins from the environment variable
	allowedOrigins := getAllowedOrigins()

	// Initialize the Gin router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize dependencies using wire
	container, err := config.InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Start the cron job
	cron := container.Cron
	defer cron.Stop()

	// Configure the application's routes
	infrastructure.ConfigureRoutes(router, container.CaptchaController, container.ParticipantController, container.VoteController, container.ResultController)

	// Initialize and start the server
	server := infrastructure.InitServer(router, port)

	// Start the vote consumer
	go container.RabbitMQConsumer.ConsumeVotes()

	// Wait for an interrupt signal to gracefully shut down the server
	infrastructure.WaitForShutdown(server)
}

// Function to retrieve the port from the environment, with a default value if not set
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// Function to retrieve the allowed origins from the environment, with a default value if not set
func getAllowedOrigins() []string {
	origins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"http://localhost:9000"}
	}
	return strings.Split(origins, ",")
}
