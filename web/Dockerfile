# Stage 1: Build the Go application
FROM golang:1.27-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o web-exe .

# Stage 2: Run the Go application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/web-exe .

# Copy static assets and views
COPY assets ./assets
COPY view ./view

# Expose the port the web server will listen on
EXPOSE 80

# Run the application
CMD ["./web-exe"]
