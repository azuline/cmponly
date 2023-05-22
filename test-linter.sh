#!/usr/bin/env bash

mkdir -p .test
# Ignore errors b/c the linter should fail.
(go run ./cmd/cmponlylint testdata/test.go >.test/test.out 2>&1) || true
# Replace filepaths for determinism.
sed -i 's/^.*test\.go:/test.go:/' .test/test.out
# Diff the output.
diff .test/test.out testdata/test.go.want
