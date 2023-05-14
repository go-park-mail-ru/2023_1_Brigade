FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o chat cmd/chat/main.go

FROM scratch

COPY --from=builder chat /bin/chat

ENTRYPOINT ["/bin/chat"]
