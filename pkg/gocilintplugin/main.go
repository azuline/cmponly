package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/azuline/cmponly/pkg/cmponlylint"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{cmponlylint.Analyzer}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin

// Have a main function so build doesn't fail.
func main() {}
