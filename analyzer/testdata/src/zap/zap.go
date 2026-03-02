package zap

import (
	"go.uber.org/zap"
)

func testZap() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	// Тесты для Logger
	logger.Info("Starting server")  // want "must start with lowercase"
	logger.Info("запуск сервера")   // want "only ASCII characters" "lowercase letters"
	logger.Info("server started!🚀") // want "only ASCII characters" "lowercase letters"

	// Тесты для SugaredLogger
	sugar.Info("Starting server")  // want "must start with lowercase"
	sugar.Info("запуск сервера")   // want "only ASCII characters" "lowercase letters"
	sugar.Info("server started!🚀") // want "only ASCII characters" "lowercase letters"

	// Тесты для методов с суффиксами
	sugar.Infow("Starting server", "port", 8080)  // want "must start with lowercase"
	sugar.Infow("server started!🚀", "port", 8080) // want "only ASCII characters" "lowercase letters"

	sugar.Infof("Starting server on %d", 8080)      // want "must start with lowercase"
	sugar.Infof("запуск сервера на порту %d", 8080) // want "only ASCII characters" "lowercase letters"

	logger.Info("starting server")               // OK
	sugar.Info("starting server")                // OK
	sugar.Infow("starting server", "port", 8080) // OK
}
