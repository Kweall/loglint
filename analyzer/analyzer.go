package analyzer

import (
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages format",
	Run:  run,
}

var configPath string

func init() {
	Analyzer.Flags.StringVar(&configPath, "config", "", "path to config file")
}
