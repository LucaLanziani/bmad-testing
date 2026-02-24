.PHONY: all build test lint clean help

# Default target
all: test lint

# Build the conciglia binary (when main.go exists)
build:
	@echo "Building conciglia..."
	@if [ -f cmd/conciglia/main.go ]; then \
		go build -o bin/conciglia ./cmd/conciglia; \
	else \
		echo "Note: cmd/conciglia/main.go not yet implemented. Skipping build."; \
	fi

# Run all tests
test:
	@echo "Running tests..."
	@go test ./internal/... -v -cover

# Run tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	@go test ./internal/... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter (golangci-lint)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Running go vet instead..."; \
		go vet ./...; \
		echo "Install golangci-lint for full linting: https://golangci-lint.run/usage/install/"; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Run goyacc (for parser generation, when grammar exists)
generate:
	@echo "Running code generation..."
	@if [ -f internal/parser/grammar.y ]; then \
		go generate ./...; \
	else \
		echo "Note: Parser grammar not yet implemented. Skipping generation."; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Run demo
demo:
	@echo "Running demo..."
	@go run examples/demo.go

# Help
help:
	@echo "Available targets:"
	@echo "  make all           - Run tests and linter (default)"
	@echo "  make build         - Build the conciglia binary"
	@echo "  make test          - Run all tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make lint          - Run linter (golangci-lint or go vet)"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo "  make generate      - Run code generation (goyacc)"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make deps          - Install/update dependencies"
	@echo "  make demo          - Run the interactive demo"
	@echo "  make help          - Show this help message"
