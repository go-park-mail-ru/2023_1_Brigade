FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o user cmd/user/main.go

FROM scratch

COPY --from=builder /app/user /

CMD ["/user"]
