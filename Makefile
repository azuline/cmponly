build:
	go build ./cmd/cmponlylint

test:
	-go run ./cmd/cmponlylint testdata/test.go >test.out 2>&1 
	diff test.out testdata/test.go.want

lint:
	golangci-lint run --fix --config .golangci.yml ./...
	go mod tidy

.PHONY: build test lint
