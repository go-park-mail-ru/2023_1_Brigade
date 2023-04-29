all: clean run

.PHONY: run
run:
	cd docker && docker compose up -d

.PHONY: clean_microservices
clean: |
	docker stop api || true
	docker stop chat || true
	docker stop user || true
	docker stop auth || true
	docker stop consumer || true
	docker stop producer || true
	docker stop messages || true

	docker rm api || true
	docker rm chat || true
	docker rm user || true
	docker rm auth || true
	docker rm consumer || true
	docker rm producer || true
	docker rm messages || true

	docker rmi docker-api || true
	docker rmi docker-chat || true
	docker rmi docker-user || true
	docker rmi docker-auth || true
	docker rmi docker-consumer || true
	docker rmi docker-producer || true
	docker rmi docker-messages || true

.PHONY: clean_images_containers
clean_images_containers: |
	docker stop $(docker ps -q)
	docker system prune -a

.PHONY: generate_proto_rpc
generate_proto_rpc: |
	protoc --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=internal/generated protobuf/*_rpc.proto

.PHONY: generate_proto
generate_proto: |
	find protobuf -type f -name '*.proto' ! -name '*_rpc.proto' -exec protoc --go_out=internal/generated {} +

.PHONY: cover_out
cover_out: |
	go test -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" | grep -v ".pb" > tmp.out
	go tool cover -func=tmp.out

.PHONY: cover_html
cover_html: |
	go test -v ./... -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" > tmp.out
	go tool cover -html=tmp.out
