include .env
export

.PHONY: build run test clean docker-up docker-down migrate

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./tests/...

# Clean build artifacts
clean:
	rm -rf bin/

# Start services with Docker Compose
docker-up:
	docker-compose up -d

# Stop services
docker-down:
	docker-compose down

# Build and run with Docker Compose
docker-build:
	docker-compose up --build

# View logs
logs:
	docker-compose logs -f api

# Install dependencies
deps:
	go mod tidy
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Generate mocks (requires mockgen)
mocks:
	go generate ./...

# Database operations
db-create:
	createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)

db-drop:
	dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)

# Help
help:
	@echo "Available commands:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  docker-up  - Start services with Docker Compose"
	@echo "  docker-down - Stop services"
	@echo "  deps       - Install dependencies"
	@echo "  fmt        - Format code"
	@echo "  lint       - Lint code"