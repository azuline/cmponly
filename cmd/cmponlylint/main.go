package main

import (
	"github.com/azuline/cmponly/pkg/cmponlylint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(cmponlylint.Analyzer)
}
