#FROM golang:1.19-alpine AS builder
#
#WORKDIR /app
#
#COPY . /app
#
#RUN go build -o auth cmd/auth/main.go
#
#FROM scratch
#
#COPY --from=builder /app/auth /
#
#CMD ["/auth"]

FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth cmd/auth/main.go

FROM scratch

COPY --from=builder auth /bin/auth

ENTRYPOINT ["/bin/auth"]

