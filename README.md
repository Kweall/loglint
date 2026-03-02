# LogLint — Линтер для проверки лог-записей в Go

**Описание:**  
LogLint — это линтер для Go, совместимый с `golangci-lint`, который проверяет лог-записи в коде на соответствие установленным правилам и предотвращает использование чувствительных данных.

## Требования

Линтер проверяет следующие правила для лог-сообщений:

1. **Сообщение должно начинаться со строчной буквы**  

❌ Неправильно:  
```go
log.Info("Starting server on port 8080")
slog.Error("Failed to connect to database")
````

✅ Правильно:

```go
log.Info("starting server on port 8080")
slog.Error("failed to connect to database")
```

2. **Сообщение должно быть только на английском языке**

❌ Неправильно:

```go
log.Info("запуск сервера")
log.Error("ошибка подключения к базе данных")
```

✅ Правильно:

```go
log.Info("starting server")
log.Error("failed to connect to database")
```

3. **Сообщение не должно содержать спецсимволы или эмодзи**

❌ Неправильно:

```go
log.Info("server started!🚀")
log.Error("connection failed!!!")
log.Warn("warning: something went wrong...")
```

✅ Правильно:

```go
log.Info("server started")
log.Error("connection failed")
log.Warn("something went wrong")
```

4. **Сообщение не должно содержать чувствительные данные**

❌ Неправильно:

```go
log.Info("user password: " + password)
log.Debug("api_key=" + apiKey)
log.Info("token: " + token)
```

✅ Правильно:

```go
log.Info("user authenticated successfully")
log.Debug("api request completed")
log.Info("token validated")
```

**Поддерживаемые логгеры:**

* `log/slog`
* `go.uber.org/zap`

**Язык:** Go 1.24.4
**Совместимость:** golangci-lint

---

## Установка и сборка

1. Клонируем репозиторий:

```bash
git clone https://github.com/kweall/loglint.git
cd loglint
```

2. Скачиваем зависимости:

```bash
go mod tidy
```

3. Собираем исполняемый бинарный файл для `go vet`:

```bash
go build -o bin/loglint ./cmd/loglint
chmod +x bin/loglint
```

---

## Использование

### 1. Проверка проекта через `go vet`:
### В корне своего проекта:
```bash
go vet -vettool=bin/loglint ./...
```

### 3. Настройка через `loglint.yaml` (опционально):

```yaml
rules:
  lowercase: true
  ascii: true
  special_chars: true
  sensitive: true

sensitive_keys:
  - password
  - token
  - api_key
  - my_custom_secret
```

* Можно включать/выключать отдельные проверки.
* Добавлять свои ключи для проверки чувствительных данных.

---

## Интеграция с `golangci-lint`

1. Установить golangci-lint на MacOS:

```bash
brew install golangci-lint
```

2. Пример `.golangci.yml` для использования LogLint:

```yaml
linters-settings:
  govet:
    check-shadowing: true

linters:
  enable:
    - govet
    - loglint

run:
  tests: true
```

3. Запуск:

```bash
golangci-lint run
```

На Windows рекомендуется использовать `go vet -vettool` напрямую, на Linux/Mac — можно интегрировать в golangci-lint через плагин.

---

## Тесты

Все тесты находятся в `analyzer/testdata` и `analyzer/`.

Запуск:

```bash
go test ./analyzer -v
```

* `TestAnalyzer` — проверка правил логов.

---

## CI/CD

Пример автоматической сборки и тестирования на GitHub Actions:

* Сборка бинаря через `go build ./cmd/loglint`
* Запуск unit-тестов
* Проверка `go vet` с вашим анализатором
* Отображение SuggestedFixes
* Проверка конфигурации `loglint.yaml`

Полный рабочий пример находится в `.github/workflows/ci.yml`.

---

## Бонусные возможности

1. **Конфигурация правил** через `loglint.yaml`.
2. **SuggestedFixes** для автоматического исправления сообщений.
3. **Кастомные паттерны** для чувствительных данных.
4. **CI/CD** для автоматической сборки и тестирования.