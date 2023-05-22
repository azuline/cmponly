package cmponly

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type TestStruct struct {
	A string
	B string
	C string
	D string
}

func TestCmpOnly(t *testing.T) {
	stc := TestStruct{A: "a", B: "b", C: "c", D: "d"}

	if diff := cmp.Diff(TestStruct{A: "a", B: "c"}, stc, Fields(TestStruct{}, "A")); diff != "" {
		t.Fatalf("(-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(TestStruct{A: "a", B: "c"}, stc, Fields(TestStruct{}, "B")); diff == "" {
		t.Fatalf("empty diff; should have found difference in B")
	}

	if diff := cmp.Diff(TestStruct{A: "a", B: "c"}, stc, Fields(TestStruct{}, "C")); diff == "" {
		t.Fatalf("empty diff; should have found difference in C")
	}
}
