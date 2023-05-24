ARG VERSION
FROM technogramm/base:$VERSION AS builder
ARG SRC_PATH
RUN go build -o service $SRC_PATH

FROM alpine
COPY --from=builder /app/service .
CMD ["./service"]