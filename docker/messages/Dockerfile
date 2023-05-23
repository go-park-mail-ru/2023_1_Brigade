#FROM golang:1.19-alpine as builder
#
#WORKDIR /app
#
#COPY go.mod go.sum /
#
#RUN go mod download
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o messages cmd/messages/main.go
#
#FROM scratch
#
#COPY --from=builder messages /bin/messages
#
#ENTRYPOINT ["/bin/messages"]

FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM builder AS build

COPY . .

RUN go build -o ./../../messages cmd/messages/main.go

CMD ["/messages"]
