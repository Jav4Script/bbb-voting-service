.PHONY: swag wire build-dev build-prod run-dev run-prod stop clean load-test clear-test-resources load-test-captcha load-test-participant load-test-vote load-test-results

# Define the TEST_FILE variable
TEST_FILE ?= load-test.js

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

# Run all load tests (including clearing test resources)
load-test: build-dev
	TEST_FILE=load-test.js $(MAKE) run-load-test

# Run only Participant-related load tests
load-test-participant: build-dev
	TEST_FILE=load-test-participant.js $(MAKE) run-load-test

# Run only CAPTCHA-related load tests
load-test-captcha: build-dev
	TEST_FILE=load-test-captcha.js $(MAKE) run-load-test

# Run only Vote-related load tests
load-test-vote: build-dev
	TEST_FILE=load-test-vote.js $(MAKE) run-load-test

# Run only Results-related load tests
load-test-results: build-dev
	TEST_FILE=load-test-results.js $(MAKE) run-load-test

# Helper target to run load tests
run-load-test:
	docker-compose -f docker-compose.test.yml up -d --force-recreate app-test test-database test-redis test-rabbitmq
	docker-compose -f docker-compose.test.yml --profile load-test up --force-recreate k6
	$(MAKE) clear-test-resources
	docker-compose -f docker-compose.test.yml down

# Stop all running containers
stop:
	docker-compose down

# Clean the project
clean:
	rm -f main
	rm -rf docs

# Clear redis data
clear-redis:
	docker exec bbb-voting-redis redis-cli FLUSHALL

# Clear test resources (Database, Redis, RabbitMQ)
clear-test-resources:
	@echo "Verify active containers..."
	docker ps -a
	@echo "Cleaning up test resources..."
	@echo "Cleaning up test database..."
	docker-compose -f docker-compose.test.yml exec test-database sh -c "chmod +x /scripts/database-cleanup.sh && /scripts/database-cleanup.sh"
	@echo "Cleaning up test Redis..."
	docker-compose -f docker-compose.test.yml exec test-redis redis-cli FLUSHALL
	@echo "Cleaning up test RabbitMQ..."
	docker-compose -f docker-compose.test.yml exec test-rabbitmq sh -c "chmod +x /scripts/rabbitmq-cleanup.sh && /scripts/rabbitmq-cleanup.sh"