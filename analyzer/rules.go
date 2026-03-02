package analyzer

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var allowed = regexp.MustCompile(`^[a-z0-9 ]+$`)

func checkLowercase(pass *analysis.Pass, call *ast.CallExpr, msg string, lit *ast.BasicLit) {
	if msg == "" {
		return
	}

	r := []rune(msg)[0]
	if unicode.IsUpper(r) {
		fixed := strings.ToLower(string(r)) + msg[1:]

		pass.Report(analysis.Diagnostic{
			Pos:     call.Pos(),
			End:     call.End(),
			Message: "log message must start with lowercase letter",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "make first letter lowercase",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(`"` + fixed + `"`),
						},
					},
				},
			},
		})
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

func checkSpecialChars(pass *analysis.Pass, call *ast.CallExpr, msg string, lit *ast.BasicLit) {
	if msg == "" {
		return
	}

	if !allowed.MatchString(msg[1:]) {

		var cleaned strings.Builder
		cleaned.WriteByte(msg[0])

		for _, r := range msg[1:] {
			if unicode.IsLower(r) || unicode.IsDigit(r) || r == ' ' {
				cleaned.WriteRune(r)
			}
		}

		pass.Report(analysis.Diagnostic{
			Pos:     call.Pos(),
			End:     call.End(),
			Message: "log message must contain only english lowercase letters, digits and spaces",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "remove special characters",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(`"` + cleaned.String() + `"`),
						},
					},
				},
			},
		})
	}
}

func checkSensitive(pass *analysis.Pass, call *ast.CallExpr, cfg *Config) {
	if len(call.Args) == 0 {
		return
	}

	switch first := call.Args[0].(type) {

	case *ast.BinaryExpr:
		if containsSensitiveInBinary(first, cfg) {
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

			if isSensitiveKey(key, cfg) {
				pass.Reportf(call.Pos(), "logging sensitive field '%s' is not allowed", key)
				return
			}
		}
	}
}

func containsSensitiveInBinary(expr ast.Expr, cfg *Config) bool {
	switch e := expr.(type) {

	case *ast.BinaryExpr:
		return containsSensitiveInBinary(e.X, cfg) ||
			containsSensitiveInBinary(e.Y, cfg)

	case *ast.BasicLit:
		val := strings.ToLower(strings.Trim(e.Value, `"`))
		return isSensitiveKey(val, cfg)

	case *ast.Ident:
		name := strings.ToLower(e.Name)
		return isSensitiveKey(name, cfg)
	}

	return false
}

func isSensitiveKey(key string, cfg *Config) bool {
	key = strings.ToLower(key)

	for _, s := range cfg.SensitiveKeys {
		if strings.Contains(key, s) {
			return true
		}
	}
	return false
}
