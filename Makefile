all: clean run_prod

.PHONY: run_local_microservices
run_local_microservices:
#	sudo kill -9 $(sudo lsof -t -i:9000) &
#	sudo kill -9 $(sudo lsof -t -i:9001) &
#	sudo kill -9 $(sudo lsof -t -i:9002) &
#	sudo kill -9 $(sudo lsof -t -i:9003) &
#	sudo kill -9 $(sudo lsof -t -i:9004) &
#	sudo kill -9 $(sudo lsof -t -i:9005) &
#	sudo kill -9 $(sudo lsof -t -i:8081)
	go run cmd/consumer/rabbitMQ/main.go >> logs/consumer 2>&1 &
	go run cmd/producer/rabbitMQ/main.go >> logs/producer 2>&1 &
	go run cmd/auth/main.go >> logs/auth 2>&1 &
	go run cmd/chat/main.go >> logs/chat 2>&1 &
	go run cmd/messages/main.go >> logs/messages 2>&1 &
	go run cmd/user/main.go >> logs/user 2>&1 &
	go run cmd/api/main.go >> logs/api 2>&1 &

.PHONY: run_stack
run_stack:
	cd docker && docker compose -f docker-compose-stack.yml up -d

.PHONY: run_prod
run_prod:
	cd docker && docker compose -f docker-compose-prod.yml up -d

.PHONY: run_sentry
run_sentry: |
	cd docker && docker compose -f docker-compose-stack.yml run --rm sentry-base config generate-secret-key
	cd docker && docker compose -f docker-compose-stack.yml run --rm sentry-base upgrade

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
	cat c.out | grep -v "cmd" | grep -v "_mock.go" | grep -v ".pb" > tmp.out
	go tool cover -html=tmp.out
