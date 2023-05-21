test:
	-go run ./cmd/cmponlylint testdata/test.go >test.out 2>&1 
	diff test.out testdata/test.go.want
