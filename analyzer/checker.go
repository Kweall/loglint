package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	cfg := loadConfig(configPath)

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isLoggerCall(call) {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			if cfg.Rules.Sensitive {
				checkSensitive(pass, call, cfg)
			}

			msgLit, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				return true
			}

			if msgLit.Kind != token.STRING {
				return true
			}

			msg := strings.Trim(msgLit.Value, `"`)

			if cfg.Rules.ASCII {
				checkEnglish(pass, call, msg)
			}

			if cfg.Rules.Lowercase {
				checkLowercase(pass, call, msg, msgLit)
			}

			if cfg.Rules.SpecialChars {
				checkSpecialChars(pass, call, msg, msgLit)
			}

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

	switch method {
	case "Infow", "Errorw", "Warnw", "Debugw":
		return true
	case "Infof", "Errorf", "Warnf", "Debugf":
		return true
	}

	return false
}
