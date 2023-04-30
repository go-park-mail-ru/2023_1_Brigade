#FROM golang:1.19-alpine as builder
#
#WORKDIR /app
#
##COPY go.mod go.sum ./
##RUN #go mod download
#
#COPY . .
#
#ENTRYPOINT ["/app/cmd/auth/auth"]
##RUN #go /app/cmd/app/auth
##RUN #go build -o /app/auth cmd/auth/main.go

# Финальный этап, копируем собранное приложение
FROM alpine:3

WORKDIR /app

COPY . .

ENTRYPOINT ["/app/cmd/auth/auth"]


#
#WORKDIR /app
#
#COPY --from=builder /app .
#
#ENTRYPOINT ["/app/auth"]
