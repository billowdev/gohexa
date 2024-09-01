#!/bin/bash

# Create the build directory if it doesn't exist
mkdir -p build

# Build for macOS (Intel/AMD64)
GOOS=darwin GOARCH=amd64 go build -o build/gohexa-mac ./cmd/main.go

# Build for macOS (Apple Silicon/ARM64)
GOOS=darwin GOARCH=arm64 go build -o build/gohexa-mac-arm ./cmd/main.go

# Build for Linux (AMD64)
GOOS=linux GOARCH=amd64 go build -o build/gohexa-linux ./cmd/main.go

# Build for Windows (AMD64)
GOOS=windows GOARCH=amd64 go build -o build/gohexa.exe ./cmd/main.go

# Build for Windows (32-bit)
GOOS=windows GOARCH=386 go build -o build/gohexa-32.exe ./cmd/main.go

echo "Builds completed for macOS, Linux, and Windows in the build/ directory."
