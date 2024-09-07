.PHONY: run init tidy build

# Dependencies
MODULES := $(wildcard cmd/**/*.go internal/**/*.go pkg/**/*.go)

# Initialize project by tidying up dependencies and running the application
init: tidy run

# Tidy up Go modules
tidy:
	go mod tidy

# Build the Go application
build:
	go build -o bin/main ./cmd/main.go

# Run the Go application
run: $(MODULES)
	go run ./cmd/main.go

test:
	go test ./... -cover