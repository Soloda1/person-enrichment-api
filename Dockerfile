# Build stage
FROM golang:1.23.6-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

RUN apk add --no-cache ca-certificates

CMD ["./main"] 