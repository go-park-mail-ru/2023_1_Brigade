FROM golang:1.19-alpine AS builder

WORKDIR /usr/local/go/src/

ADD . /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o chat cmd/chat/main.go

VOLUME /usr/local/go/src/avatars /home/ubuntu/2023_1_Brigade/avatars

FROM alpine:3.14

COPY --from=builder /usr/local/go/src/ /
COPY --from=builder /usr/local/go/src/internal/configs/config.yaml /
COPY --from=builder /usr/local/go/src/cmd/api/background.png /
COPY --from=builder /usr/local/go/src/avatars /

CMD ["/chat"]
