FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM builder AS build

COPY . .

RUN go build -o ./../../messages cmd/messages/main.go

CMD ["/messages"]
