FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/consumer cmd/consumer/main.go

# Финальный этап, копируем собранное приложение
FROM alpine:3

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["/app/consumer"]
