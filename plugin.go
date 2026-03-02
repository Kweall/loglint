package main

import (
	"github.com/kweall/loglint/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}
