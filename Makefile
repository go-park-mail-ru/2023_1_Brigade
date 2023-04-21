PROTO_FILES = $(shell find . -iname '*.proto')

all: run clean

.PHONY: run
run: ## Run project
	docker compose up -d

.PHONY: stop
clean: ## Clean containers and images
	docker compose kill
	docker compose down

.PHONY: test
test: ## Run all the tests
	go test ./...

.PHONY: proto

proto:
	protoc -I=protobuf --go_out=plugins=grpc:client protobuf/chat.proto
#    protoc -I=protobuf --go_out=plugins=grpc:client_generated protobuf/chat.proto
#    protoc -I=protobuf --grpc-gateway_out=logtostderr=true,paths=source_relative,grpc_api_configuration=protobuf/chat.yaml:client_generated protobuf/chat.proto

#.PHONY: proto
#proto: ## Make protobuf files
#	protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_FILES)

.PHONY: cover_out
cover_out: test ## Run all the tests and opens the coverage report
	go test -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" > tmp.out
	go tool cover -func=tmp.out

.PHONY: cover_html
cover_html: test ## Run all the tests and opens the coverage report in HTML
	go test -v ./... -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" > tmp.out
	go tool cover -html=tmp.out
