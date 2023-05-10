FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o auth cmd/auth/main.go

FROM scratch

COPY --from=builder /app/auth /

CMD ["/auth"]
