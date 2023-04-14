FROM golang:1.19-alpine AS builder

WORKDIR /usr/local/go/src/

ADD . /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/api/main.go

FROM alpine:3.14

COPY --from=builder /usr/local/go/src/ /
COPY --from=builder /usr/local/go/src/internal/configs/config.yaml /

CMD ["/app"]
#
#FROM redis:latest
#COPY redis.conf /usr/local/etc/redis/redis.conf
#CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
