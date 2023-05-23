#FROM golang:1.19-alpine AS builder
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#
#RUN go mod download
#
#FROM builder AS build
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o ./../../auth cmd/auth/main.go
#
#CMD ["/auth"]

FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM builder AS build

COPY . .

RUN go build -o ./../../api cmd/api/main.go

CMD ["/api"]
