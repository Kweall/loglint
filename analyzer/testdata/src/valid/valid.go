package valid

import "log/slog"

func test() {
	// lowercase
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	// english
	slog.Info("starting server")
	slog.Error("failed to connect to database")
	// special
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
	//sensitive
	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")
}
