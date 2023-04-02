FROM golang:1.19 AS builder

WORKDIR /usr/local/go/src/

ADD . /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8001

COPY --from=builder /usr/local/go/src/ /
COPY --from=builder /usr/local/go/src/internal/configs/config.yaml /

CMD ["/app"]
