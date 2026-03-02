package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "the log linter",
	Run:  run,
}
