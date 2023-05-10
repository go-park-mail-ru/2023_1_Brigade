FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN go build -o api cmd/api/main.go

FROM scratch

COPY --from=builder api /bin/api

ENTRYPOINT ["/bin/api"]
