FROM golang:1.18

WORKDIR /app
COPY . /app
COPY internal/configs/config.yaml /app
COPY .env /app
RUN go build cmd/api/main.go

EXPOSE 8081
CMD ["./main"]
