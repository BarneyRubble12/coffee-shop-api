.PHONY: all build run clean test swagger deps help dev install-tools debug

# Binary name
BINARY_NAME=coffee-shop-api

# Build flags
LDFLAGS=-ldflags "-w -s"

# Default target
all: clean build

# Build the application
build:
	@echo "Building..."
	@go build $(LDFLAGS) -o bin/$(BINARY_NAME) cmd/coffee-shop-api/main.go

# Run the application
run:
	@echo "Running..."
	@go run cmd/coffee-shop-api/main.go

# Clean build files
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Generate Swagger documentation
swagger:
	@echo "Installing swag..."
	@go install github.com/swaggo/swag/cmd/swag@v1.8.12
	@echo "Generating Swagger docs..."
	@~/go/bin/swag init -g cmd/coffee-shop-api/main.go
	@mv docs/* internal/docs/ 2>/dev/null || true
	@rm -rf docs

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Install development tools (run manually if needed)
install-tools:
	@echo "Installing development tools..."
	@go install github.com/cosmtrek/air@v1.40.4
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/go-delve/delve/cmd/dlv@latest

# Development workflow (fast, does not install tools)
dev: deps swagger
	@echo "Starting development server..."
	@~/go/bin/air

# Debug mode (fast, does not install tools)
debug: deps swagger
	@echo "Starting development server in debug mode..."
	@dlv debug cmd/coffee-shop-api/main.go --headless --listen=:2345 --api-version=2 --accept-multiclient

# Help command
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make clean         - Clean build files"
	@echo "  make test          - Run tests"
	@echo "  make swagger       - Generate Swagger documentation"
	@echo "  make deps          - Install dependencies"
	@echo "  make all           - Clean and build the application"
	@echo "  make install-tools - Install development tools (air, swag, dlv)"
	@echo "  make dev           - Start development server with hot-reload"
	@echo "  make debug         - Start development server in debug mode" 