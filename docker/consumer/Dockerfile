FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o consumer cmd/consumer/rabbitMQ/main.go

FROM scratch

COPY --from=builder /app/consumer /

CMD ["/consumer"]
