FROM golang:1.18

WORKDIR /app
COPY . .
COPY cmd/configs/config.yaml /app
RUN go build cmd/api/main.go

EXPOSE 8082
CMD ["./main"]
