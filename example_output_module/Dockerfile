# Use the official golang image as base
FROM golang:1.18 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM debian:buster-slim

# Set the current working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/main .

# Copy the prisma folder
COPY prisma/ ./prisma/

# Expose port 8080 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]
