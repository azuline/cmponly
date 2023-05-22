package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/azuline/cmponly/cmponlylint"
)

func main() {
	singlechecker.Main(cmponlylint.Analyzer)
}
