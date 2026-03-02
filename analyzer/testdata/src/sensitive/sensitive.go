package sensitive

import "log/slog"

func test() {
	apiKey := "askcbnOAJXAJBXCSSOmc"
	password := "pass123"
	token := "asdAbn2b1"
	slog.Info("user password: " + password) // want "sensitive data"
	slog.Debug("api_key=" + apiKey)         // want "sensitive data"
	slog.Info("token: " + token)            // want "sensitive data"
}
