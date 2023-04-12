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
