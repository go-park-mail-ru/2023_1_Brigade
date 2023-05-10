FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o producer cmd/producer/rabbitMQ/main.go

FROM scratch

COPY --from=builder /app/producer /

CMD ["/producer"]
