package main

import (
	"github.com/kweall/loglint/analyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(analyzer.Analyzer)
}
