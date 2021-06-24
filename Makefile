.PHONY: help
help: ## Shows all make targets available
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/##//'

.PHONY: serve
serve: ## run a docker container with cronlog added to the path environment variable.
	@docker build -t conlog .
	@docker run -it conlog

.PHONY: test
test: ## Test cli.
	@go test -race ./...

.PHONY: cover
cover: ## Cover cli.
	@go test -coverprofile=cover.out ./...

.PHONY: lint
lint: ## Lint cli.
	@golangci-lint run