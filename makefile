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
run-dev: build-dev
	docker-compose up --force-recreate

# Execute the application for production
run-prod: swag wire
	docker run -p 8080:8080 --env-file .env bbb-voting-service-prod

# Run load tests
load-tests: build-dev
	docker-compose -f docker-compose.test.yml up -d --force-recreate app-test test-database test-redis test-rabbitmq
	docker-compose -f docker-compose.test.yml --profile load-test up --force-recreate k6
	clear-test-resources
	docker-compose -f docker-compose.test.yml down


# Stop all running containers
stop:
	docker-compose down

# Clean the project
clean:
	rm -f main
	rm -rf docs

# Clear redis
clear-redis:
	docker exec -it bbb-voting-redis redis-cli FLUSHALL

# Clear test resources
clear-test-resources:
	docker-compose -f docker-compose.test.yml exec -it test-database sh -c "chmod +x /scripts/database-cleanup.sh && /scripts/database-cleanup.sh"
	docker-compose -f docker-compose.test.yml exec -it test-redis redis-cli FLUSHALL
	docker-compose -f docker-compose.test.yml exec -it test-rabbitmq sh -c "chmod +x /scripts/rabbitmq-cleanup.sh && /scripts/rabbitmq-cleanup.sh"