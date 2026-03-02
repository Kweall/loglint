package english

import "log/slog"

func test() {
	slog.Info("запуск сервера")                    // want "ASCII characters" "english lowercase"
	slog.Error("ошибка подключения к базе данных") // want "ASCII characters" "english lowercase"
}
