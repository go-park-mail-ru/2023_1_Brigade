.PHONY: clean_local_microservices
clean_local_microservices:
	sudo kill -9 $(sudo lsof -t -i:9000) &
	sudo kill -9 $(sudo lsof -t -i:9001) &
	sudo kill -9 $(sudo lsof -t -i:9002) &
	sudo kill -9 $(sudo lsof -t -i:9003) &
	sudo kill -9 $(sudo lsof -t -i:9004) &
	sudo kill -9 $(sudo lsof -t -i:9005) &
	sudo kill -9 $(sudo lsof -t -i:8081) &

.PHONY: run_local_microservices
run_local_microservices:
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

.PHONY: stop_prod
stop_prod:
	cd docker && docker compose -f docker-compose-prod.yml down

.PHONY: run_sentry
run_sentry: |
	cd docker && docker compose -f docker-compose-stack.yml run --rm sentry-base config generate-secret-key
	cd docker && docker compose -f docker-compose-stack.yml run --rm sentry-base upgrade

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

.PHONY: generate_easyjson
generate_easyjson: |
	cd internal/model && easyjson --all *.go

.PHONY: cover_out
cover_out: |
	go test -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "_mock.go" | grep -v ".pb" | grep -v "_easyjson.go" > tmp.out
	go tool cover -func=tmp.out

.PHONY: cover_html
cover_html: |
	go test -v ./... -coverprofile=c.out ./... -coverpkg=./...
	cat c.out | grep -v "cmd" | grep -v "_mock.go" | grep -v ".pb" | grep -v "_easyjson.go" > tmp.out
	go tool cover -html=tmp.out
