FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM builder AS build

COPY . .
COPY internal/configs/config.yaml /
COPY cmd/api/background.png /

RUN go build -o ./../../api cmd/api/main.go

CMD ["/api"]
