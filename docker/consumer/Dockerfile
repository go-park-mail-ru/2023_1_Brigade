#FROM golang:1.19-alpine AS builder
#
#WORKDIR /app
#
#COPY . /app
#
#RUN go build -o consumer cmd/consumer/rabbitMQ/main.go
#
#FROM scratch
#
#COPY --from=builder /app/consumer /
#
#CMD ["/consumer"]

FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN go build -o consumer cmd/consumer/rabbitMQ/main.go

FROM scratch

COPY --from=builder consumer /bin/consumer

ENTRYPOINT ["/bin/consumer"]
