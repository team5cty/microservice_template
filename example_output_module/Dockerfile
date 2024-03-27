# Use the official golang image as base
FROM golang:1.22.0 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application with debugging symbols
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o main .

# Start a new stage from scratch
FROM debian:buster-slim

# Set the current working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
