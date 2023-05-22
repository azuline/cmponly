package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/azuline/cmponly/pkg/cmponlylint"
)

func main() {
	singlechecker.Main(cmponlylint.Analyzer)
}
