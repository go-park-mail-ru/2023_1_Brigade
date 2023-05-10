#FROM golang:1.19-alpine AS builder
#
#WORKDIR /app
#
#COPY . /app
#
#RUN go build -o user cmd/user/main.go
#
#FROM scratch
#
#COPY --from=builder /app/user /
#
#CMD ["/user"]

FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN go build -o user cmd/user/main.go

FROM scratch

COPY --from=builder user /bin/user

ENTRYPOINT ["/bin/user"]
