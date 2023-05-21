package test

import (
	"testing"

	"github.com/azuline/cmponly/pkg/cmponly"
	"github.com/google/go-cmp/cmp"
)

type TestStruct struct {
	A string
	B string
	C string
	D string
}

func Example(t *testing.T) {
	stc := TestStruct{A: "a", B: "b", C: "c", D: "d"}

	// Non-error use cases.
	cmp.Diff(stc, stc, cmponly.Fields(stc, "A", "B"))
	cmp.Diff(stc, stc, cmponly.Fields(TestStruct{}, "A", "B"))
	cmp.Diff(stc, stc, cmponly.Fields(stc))
	cmp.Diff(stc, stc, cmponly.Fields(TestStruct{}))

	// Error use cases.
	cmp.Diff(stc, stc, cmponly.Fields(stc, "E"))
	cmp.Diff(stc, stc, cmponly.Fields(TestStruct{}, "E"))
	cmp.Diff(stc, stc, cmponly.Fields(stc, "A", "E"))
	cmp.Diff(stc, stc, cmponly.Fields(TestStruct{}, "A", "E"))
}
