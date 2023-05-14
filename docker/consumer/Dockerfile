FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o consumer cmd/consumer/rabbitMQ/main.go

FROM scratch

COPY --from=builder consumer /bin/consumer

ENTRYPOINT ["/bin/consumer"]
