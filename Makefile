.PHONY: help build fmt
.DEFAULT_GOAL := help

build: ## Build binary
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o gocovmerge main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

fmt: ## Format source code
	go fmt ./...
