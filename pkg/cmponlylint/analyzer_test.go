package cmponlylint

import (
	"fmt"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	testdata := analysistest.TestData()
	fmt.Printf("%+v\n", testdata)
	analysistest.Run(t, testdata, Analyzer, "cmponly")
}
