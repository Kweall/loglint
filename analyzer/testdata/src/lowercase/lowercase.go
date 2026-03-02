package lowercase

import "log/slog"

func test() {
	slog.Info("Starting server")                // want "must start with lowercase"
	slog.Info("Starting server on port 8080")   // want "must start with lowercase"
	slog.Error("Failed to connect to database") // want "must start with lowercase"
}
