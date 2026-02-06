# Builder stage
FROM golang:1.25.6-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .
COPY wait-for-db.sh .

# Default command (can be overridden by docker-compose)
CMD ["./main"]
