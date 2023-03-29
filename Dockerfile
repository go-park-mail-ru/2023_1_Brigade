FROM golang:1.18-alpine3.14 AS builder

WORKDIR /usr/local/go/src/

ADD . /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/api/main.go

FROM alpine:3.14

COPY --from=builder /usr/local/go/src/ /
COPY --from=builder /usr/local/go/src/internal/configs/config.yaml /

CMD ["/app"]
