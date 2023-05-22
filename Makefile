build:
	go build ./cmd/cmponlylint

test:
	go test ./...
	./test-lint.sh

lint:
	golangci-lint run --fix --config .golangci.yml ./...
	go mod tidy

.PHONY: build test lint
