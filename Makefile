build:
	go build ./cmd/cmponlylint

test:
	go test ./...
	./test-linter.sh

lint:
	golangci-lint run --fix --config .golangci.yml ./...
	go mod tidy

.PHONY: build test lint
