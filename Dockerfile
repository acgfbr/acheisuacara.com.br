# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git for private repos if needed
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Add ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/main .

# Create directory for environment file
RUN mkdir -p /app/config

# Copy environment file if it exists, otherwise use example
COPY .env* /app/config/
RUN if [ ! -f /app/config/.env ]; then \
    cp /app/config/.env.example /app/config/.env; \
    fi

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"] 