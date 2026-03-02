package special

import "log/slog"

func test() {
	slog.Info("server started!🚀")                 // want "only ASCII characters" "lowercase letters"
	slog.Error("connection failed!!!")            // want "lowercase letters"
	slog.Warn("warning: something went wrong...") // want "lowercase letters"
}
