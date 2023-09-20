#ARG VERSION
#FROM technogramm/base:$VERSION AS builder
#ARG SRC_PATH
#RUN go build -o service $SRC_PATH
#
#FROM alpine
#COPY --from=builder /app/service /app/.env /app/internal/config/config.yaml ./
#CMD ["./service"]

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
COPY /app/internal/config/config.yaml ./

RUN go build -o ./../../api cmd/api/main.go

CMD ["/api"]
