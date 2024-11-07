# Dockerfile

FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
COPY config/config.yaml /app/config/config.yaml
COPY pkg/db/migrations /app/migrations

EXPOSE 8080
CMD ["./main"]
