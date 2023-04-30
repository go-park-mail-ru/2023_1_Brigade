#FROM golang:1.19-alpine as builder
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .
#
#RUN go build -o /app/producer cmd/producer/main.go
#
## Финальный этап, копируем собранное приложение
#FROM alpine:3
#
#WORKDIR /app
#
#COPY --from=builder /app .
#
#ENTRYPOINT ["/app/producer"]












FROM alpine:3

WORKDIR /app

COPY . .

ENTRYPOINT ["/app/cmd/producer/producer"]

