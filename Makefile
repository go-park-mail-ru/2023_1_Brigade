PROTO_FILES = $(shell find . -iname '*.proto')
ACTIVE_PACKAGES = $(shell go list ./... | grep -Ev "mocks|generated" | tr '\n' ',')

all: clean run

.PHONY: run
run:
	cd docker && docker compose up -d

.PHONY: clean
clean: |
	docker stop api || true
	docker stop chat || true
	docker stop user || true
	docker stop messages || true
	docker rm api || true
	docker rm chat || true
	docker rm user || true
	docker rm messages || true
	docker rmi docker-api || true
	docker rmi docker-chat || true
	docker rmi docker-user || true
	docker rmi docker-messages || true

.PHONY: test
test:
	go test ./...

.PHONY: proto
proto:
	protoc -I=protobuf --go_out=plugins=grpc:client protobuf/chat.proto

#.PHONY: cover_out
#cover_out: test
	#go test -coverpkg=$(ACTIVE_PACKAGES) -coverprofile=c.out ./...
	#cat c.out | grep -v "cmd" | grep -v "easyjson" > tmp.out
	#go tool cover -func=tmp.out

.PHONY: cover_out
cover_out: test ## Run all the tests and opens the coverage report
	go test -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" | grep -v ".pb" > tmp.out
	go tool cover -func=tmp.out

.PHONY: cover_html
cover_html: test ## Run all the tests and opens the coverage report in HTML
	go test -v ./... -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" > tmp.out
	go tool cover -html=tmp.out
