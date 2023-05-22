# cmponly

Partial struct equality assertions in Go.

`cmponly` provides a `cmponly.Fields` function, which can be used as an
`Option` in `go-cmp`'s `cmp.Diff`. `cmponly.Fields` allows specifying a subset
of struct fields to compare. `cmponly.Fields` is the inverse of
`cmpopts.IgnoreFields`.

So for example:

```go
import (
	"testing"
	"github.com/azuline/cmponly/pkg/cmponly"
	"github.com/google/go-cmp/cmp"
)

type Record struct {
  A string
  B string
}

func ExampleTest(t *testing.T) {
  want := Record{A: "aaaa"}
  got := functionUnderTest()

  if diff := cmp.Diff(want, got, cmponly.Fields(Record{}, "A")); diff != "" {
    t.Fatalf("(-want +got):\n%s", diff)
  }
}
```

However, this pattern comes with a footgun: if we refactor `Record`'s fields to
be named `C` and `D` instead, `cmponly.Fields` will not autorefactor or fail
typechecking. `cmp.Diff` will compare 0 fields and vacuously pass.

The `cmponlylint` linter solves this footgun. This linter checks that all
specified fields are valid fields on the struct being compared. So in the above
example, the linter would fail with:

```go
example.go:9:32 specified fields do not exist on struct: A
```

## Motivation & Tradeoffs

brandur's [PartialEqual](https://brandur.org/fragments/partial-equal) describes
the motivation well.

After playing around with brandur's `PartialEqual`, I thought the zero-values
footgun too dangerous. `cmponly` is an alternative with different tradeoffs.

I find the tradeoffs in `cmponly` to be more palatable, as `cmponly`:

1. Supports comparing zero-values.
2. Is refactoring semi-friendly, as the included linter catches invalid fields.

However, this package has a few cons:

- Fields must be specified twice, once in the `want` struct, and again in the
  `cmponly.Fields` function call.
- It is refactoring semi-unfriendly, does not support LSP auto-rename.
- It requires an additional linter.

## Usage

TODO: standalone

TODO: golangci-lint

## Development

A Nix dev shell is provided to configure the development environment.

Tests can be ran with `make test`. Since `analysistest` does not support Go
modules, we have instead implemented a hacky test harness that checks the
output of the self-contained linter command.

## License

```
Copyright 2023 blissful <blissful@sunsetglow.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
