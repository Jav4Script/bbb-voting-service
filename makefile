.PHONY: swag wire build-dev build-prod run-dev run-prod stop clean

# Generate swagger docs
swag:
	swag init -g cmd/main.go

# Generate wire dependencies
wire:
	wire ./internal/infrastructure/config

# Compile the application for development
build-dev: swag wire
	docker-compose build

# Compile the application for production
build-prod: swag wire
	docker build -f Dockerfile.prod -t bbb-voting-service-prod .

# Execute the application for development
run-dev: swag wire
	docker-compose up --force-recreate

# Execute the application for production
run-prod: swag wire
	docker run -p 8080:8080 --env-file .env bbb-voting-service-prod

# Stop all running containers
stop:
	docker-compose down

# Clean the project
clean:
	rm -f main
	rm -rf docs