# --- Stage 1: Build the Go binary ---
FROM golang:1.24-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the Go binary
RUN go build -o main .

# --- Stage 2: Create minimal runtime image ---
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port (adjust if your server uses another)
EXPOSE 8080

# Run the binary
CMD ["./main"]
