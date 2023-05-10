FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o messages cmd/messages/main.go

FROM scratch

COPY --from=builder /app/messages /

CMD ["/messages"]
