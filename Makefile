# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building for the current platform..."
	@go build -o build/secretsharing ./cmd/secretsharing

# Cross-platform builds
#Build for Windows
build-windows:
	@echo "Building for Windows..."
	set GOOS=windows && set GOARCH=amd64 && go build -o build/windows/secretsharing.exe ./cmd/secretsharing

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	set GOOS=linux && set GOARCH=amd64 && go build -o build/linux/secretsharing ./cmd/secretsharing

# Build for macOS
build-darwin:
	@echo "Building for macOS..."
	set GOOS=darwin && set GOARCH=amd64 && go build -o build/darwin/secretsharing ./cmd/secretsharing

# Run the application
run:
	@go run cmd/secretsharing/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integration Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


.PHONY: all build run test clean watch itest
