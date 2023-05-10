FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o chat cmd/chat/main.go

FROM scratch

COPY --from=builder /app/chat /

CMD ["/chat"]
