FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM builder AS build

COPY . .
COPY /app/internal/config/config.yaml ./

RUN go build -o api cmd/api/main.go

CMD ["/api"]
