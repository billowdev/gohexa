# Start from the official Go image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code and templates into the container
COPY . .

# Build the Go binary from the specified source file
RUN go build -o gohexa ./cmd/main.go

# Create a smaller image for running the binary
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gohexa /usr/local/bin/gohexa

# Copy the templates directory from the builder stage
COPY --from=builder /app/templates /app/templates

# Set the entrypoint for the container
ENTRYPOINT ["gohexa"]

# Default command (can be overridden)
CMD ["--help"]
