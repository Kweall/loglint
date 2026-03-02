package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			if !isLoggerCall(call) {
				return true
			}

			msg := extractMessage(call)

			checkRules(pass, call, msg)
			return true
		})
	}

	return nil, nil
}

func isLoggerCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	method := sel.Sel.Name
	switch method {
	case "Info", "Error", "Warn", "Debug":
		return true
	}

	return false
}

func extractMessage(call *ast.CallExpr) string {
	if len(call.Args) == 0 {
		return ""
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return ""
	}

	return strings.Trim(lit.Value, `"`)
}
