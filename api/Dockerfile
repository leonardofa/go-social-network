# Stage 1: Build the Go application
FROM golang:1.26-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api-exe .

# Stage 2: Run the Go application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api-exe .

# Expose the port the API will listen on (default 5000)
EXPOSE 5000

# Run the application
CMD ["./api-exe"]
