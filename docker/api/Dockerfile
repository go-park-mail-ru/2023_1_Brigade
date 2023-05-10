FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o api cmd/api/main.go

FROM scratch

COPY --from=builder /app/api /

CMD ["/api"]
