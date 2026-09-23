# Stage 1: Build binary
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

# Copy dependency definitions
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final lightweight image
FROM alpine:latest

WORKDIR /app

# Copy binary and .env file from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 3000

CMD ["./main"]
