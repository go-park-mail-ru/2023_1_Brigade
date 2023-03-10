.PHONY: test
test: ## Run all the tests
	go test ./...

.PHONY: cover_out
cover_out: test ## Run all the tests and opens the coverage report
	go test -coverprofile=c.out ./...
	cat c.out | grep -v "cmd" | grep -v "easyjson" > tmp.out
	go tool cover -func=tmp.out

.PHONY: cover_html
cover_html: test ## Run all the tests and opens the coverage report in HTML
	go test -v ./... -coverprofile=c.out ./...
	cat c.out | grep -v "cmd" | grep -v "easyjson" > tmp.out
	go tool cover -html=tmp.out
