# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod
COPY go.mod ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy config directory
COPY --from=builder /app/config ./config

# Create uploads directory
RUN mkdir -p uploads

# Expose port (Backend runs on 8081)
EXPOSE 8081

# Command to run
CMD ["./main"]
