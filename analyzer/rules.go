package analyzer

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var allowed = regexp.MustCompile(`^[a-z0-9 ]+$`)

var sensitiveKeys = []string{
	"password",
	"passwd",
	"pwd",
	"token",
	"api_key",
	"apikey",
	"secret",
	"private_key",
	"auth",
}

func checkRules(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	checkEnglish(pass, call, msg)
	checkLowercase(pass, call, msg)
	checkSpecialChars(pass, call, msg)
	checkSensitive(pass, call)
}

func checkLowercase(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if msg == "" {
		return
	}

	r := []rune(msg)[0]
	if !unicode.IsLower(r) {
		pass.Reportf(call.Pos(), "log message must start with lowercase letter")
	}
}

func checkEnglish(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII {
			pass.Reportf(call.Pos(), "log message must contain only ASCII characters")
			return
		}
	}
}

func checkSpecialChars(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if msg == "" {
		return
	}

	if !allowed.MatchString(msg[1:]) {
		pass.Reportf(call.Pos(), "log message must contain only english lowercase letters, digits and spaces")
	}
}

func checkSensitive(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	switch first := call.Args[0].(type) {

	case *ast.BinaryExpr:
		if containsSensitiveInBinary(first) {
			pass.Reportf(call.Pos(), "log message may contain sensitive data")
			return
		}

	case *ast.CallExpr:
		pass.Reportf(call.Pos(), "log message must not use dynamic formatting with sensitive data")
		return
	}

	if len(call.Args) > 1 {
		for i := 1; i < len(call.Args); i += 2 {

			keyLit, ok := call.Args[i].(*ast.BasicLit)
			if !ok {
				continue
			}

			key := strings.ToLower(strings.Trim(keyLit.Value, `"`))

			if isSensitiveKey(key) {
				pass.Reportf(call.Pos(), "logging sensitive field '%s' is not allowed", key)
				return
			}
		}
	}
}

func containsSensitiveInBinary(expr ast.Expr) bool {
	switch e := expr.(type) {

	case *ast.BinaryExpr:
		return containsSensitiveInBinary(e.X) || containsSensitiveInBinary(e.Y)

	case *ast.BasicLit:
		val := strings.ToLower(strings.Trim(e.Value, `"`))
		return isSensitiveKey(val)

	case *ast.Ident:
		name := strings.ToLower(e.Name)
		return isSensitiveKey(name)
	}

	return false
}

func isSensitiveKey(key string) bool {
	for _, s := range sensitiveKeys {
		if strings.Contains(key, s) {
			return true
		}
	}
	return false
}
