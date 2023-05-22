build:
	go build ./cmd/cmponlylint

test:
	-go run ./cmd/cmponlylint testdata/test.go >test.out 2>&1 
	# Replace filepaths for determinism.
	sed -i 's/^.*test\.go:/test.go:/' test.out
	diff test.out testdata/test.go.want

lint:
	golangci-lint run --fix --config .golangci.yml ./...
	go mod tidy

.PHONY: build test lint
