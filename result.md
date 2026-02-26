# Анализ проекта hexlet-go-crawler

## Обзор проекта

**hexlet-go-crawler** — утилита командной строки для анализа структуры веб-сайтов с проверкой битых ссылок и SEO-параметров.

## Соответствие требованиям ТЗ

### ✅ Инициализация и структура проекта

| Требование | Статус | Примечание |
|------------|--------|------------|
| Модуль называется `code` | ✅ | `go.mod: module code` |
| Функция `Analyze()` в пакете `code/crawler` | ✅ | `internal/crawler/crawler.go` |
| Утилита доступна через `go run ./cmd/hexlet-go-crawler` | ✅ | Реализовано |
| CLI на базе `urfave/cli/v3` | ✅ | Реализовано |

### ✅ Makefile команды

| Команда | Статус | Описание |
|---------|--------|----------|
| `make build` | ✅ | Компилирует в `bin/hexlet-go-crawler` |
| `make test` | ✅ | Запускает `go test ./...` |
| `make run URL=<адрес>` | ✅ | Запускает утилиту с URL |
| `make lint` | ✅ | Запускает `golangci-lint` |
| `make clean` | ✅ | Удаляет `bin/` |

### ✅ Тестируемость HTTP-кода

| Требование | Статус | Примечание |
|------------|--------|------------|
| HTTP через `*http.Client` | ✅ | Передается в `Options` |
| Запрет `http.DefaultClient` / `http.Get` | ✅ | Не используется |
| Подмена клиента в тестах | ✅ | `server.Client()` в тестах |
| Детерминированные тесты | ✅ | Используют `httptest.Server` |

## Архитектура проекта

```
code/
├── cmd/hexlet-go-crawler/
│   └── main.go              # Точка входа, CLI на urfave/cli/v3
├── internal/
│   ├── crawler/
│   │   ├── crawler.go       # Основная логика: Analyze(), воркеры
│   │   ├── options.go       # Структуры: Options, Report, PageReport
│   │   └── crawler_test.go  # Unit-тесты (1000+ строк)
│   └── shared/
│       ├── rate_limiter.go  # Ограничитель RPS/delay
│       ├── rate_limiter_test.go
│       └── shared.go
├── go.mod
├── Makefile
└── README.md
```

## Ключевые компоненты

### 1. CLI (`cmd/hexlet-go-crawler/main.go`)

Использует `github.com/urfave/cli/v3` с флагами:
- `--depth` (int, default: 10)
- `--retries` (int, default: 1)
- `--delay` (duration)
- `--timeout` (duration, default: 15s)
- `--user-agent` (string)
- `--workers` (int, default: 4)
- `--indent` (bool, default: true)
- `--rps` (int)

### 2. Crawler (`internal/crawler/crawler.go`)

**Функция `Analyze(ctx, opts)`:**
- Создает пул воркеров (`Concurrency`)
- Использует каналы для распределения задач
- Отслеживает посещенные URL (дедупликация)
- Применяет rate limiting
- Обходит только в пределах домена

**Воркеры:**
```go
func worker(ctx, opts, rootHost, limiter, jobChan, resultChan)
```
- Берут задачи из канала
- Применяют rate limiter
- Вызывают `analyzePage()`
- Отправляют результаты в канал

**Анализ страницы:**
- HTTP-запрос через `opts.HTTPClient`
- Парсинг HTML для SEO-тегов
- Проверка ссылок на битые (параллельно)

### 3. Rate Limiter (`internal/shared/rate_limiter.go`)

Поддерживает два режима:
- `delay` — фиксированная задержка через `time.Ticker`
- `rps` — конвертируется в delay (`time.Second / rps`)

### 4. Структуры данных (`options.go`)

```go
type Report struct {
    RootURL     string
    Depth       int
    GeneratedAt string
    Pages       []PageReport
}

type PageReport struct {
    URL          string
    Depth        int
    HTTPStatus   int
    Status       string
    Error        string
    BrokenLinks  []BrokenLink
    DiscoveredAt string
    SEO          *SEOReport
    RawBody      []byte `json:"-"`
}

type SEOReport struct {
    HasTitle       bool
    Title          string
    HasDescription bool
    Description    string
    HasH1          bool
}
```

## Горутины и конкурентность

### Места использования горутин

| Компонент | Описание |
|-----------|----------|
| **Пул воркеров** | `Concurrency` горутин обрабатывают URL параллельно |
| **Проверка ссылок** | `checkLinks()` запускает по горутине на каждую ссылку |
| **Rate Limiter** | Использует `time.Ticker` в отдельной горутине (неявно) |

### Синхронизация

| Механизм | Где используется |
|----------|------------------|
| `sync.WaitGroup` | Ожидание завершения воркеров |
| `sync.Mutex` | Защита `visited` map от race condition |
| Каналы (`jobChan`, `resultChan`) | Передача задач и результатов |
| `context.Context` | Отмена обхода, таймауты |
| Semaphore (channel) | Ограничение параллелизма в `checkLinks()` |

### Поток данных

```
┌─────────────┐
│  main.go    │
└──────┬──────┘
       │ Analyze(ctx, opts)
       ▼
┌─────────────────────────────────────┐
│  crawler.Analyze()                  │
│  - создает jobChan, resultChan      │
│  - запускает N воркеров             │
│  - отправляет первый job            │
└──────────────┬──────────────────────┘
               │
    ┌──────────┴──────────┐
    │                     │
    ▼                     ▼
┌─────────┐         ┌──────────┐
│ jobChan │         │resultChan│
└────┬────┘         └────┬─────┘
     │                   │
     ▼                   ▼
┌─────────────┐    ┌─────────────┐
│   worker    │    │  сборщик    │
│ (goroutine) │    │  результатов│
└──────┬──────┘    └─────────────┘
       │
       ▼
┌──────────────┐
│ analyzePage()│
│ - HTTP запрос│
│ - parseSEO() │
│ - checkLinks()│
└──────────────┘
```

## Тестовое покрытие

### Тесты в `crawler_test.go`

| Категория | Тесты |
|-----------|-------|
| **Базовые** | `TestAnalyze_Success`, `TestAnalyze_NetworkError`, `TestAnalyze_NonOKStatus`, `TestAnalyze_ServerError`, `TestAnalyze_InvalidURL` |
| **Context** | `TestAnalyze_ContextCancellation`, `TestAnalyze_ContextCancellationPartialReport` |
| **JSON** | `TestAnalyze_IndentJSON` |
| **Broken Links** | `TestAnalyze_BrokenLinks`, `TestAnalyze_BrokenLinks_NetworkError`, `TestAnalyze_IgnoresInvalidSchemes`, `TestAnalyze_ResolvedRelativeLinks` |
| **SEO** | `TestAnalyze_SEO_AllTags`, `TestAnalyze_SEO_NoTags`, `TestAnalyze_SEO_HTMLEntities` |
| **Depth** | `TestAnalyze_DepthZero`, `TestAnalyze_DepthOne`, `TestAnalyze_ExternalLinksIgnored` |
| **Дедупликация** | `TestAnalyze_DuplicateLinks` |

### Покрытие сценариев ТЗ

| Требование | Тест |
|------------|------|
| Успешный запрос (200 OK) | `TestAnalyze_Success` |
| Ошибка сети | `TestAnalyze_NetworkError` |
| 404/500 статусы | `TestAnalyze_NonOKStatus`, `TestAnalyze_ServerError` |
| Таймауты | `TestAnalyze_NetworkError` (100ms timeout) |
| Битые ссылки (4xx/5xx) | `TestAnalyze_BrokenLinks` |
| Битые ссылки (сеть) | `TestAnalyze_BrokenLinks_NetworkError` |
| SEO теги присутствуют | `TestAnalyze_SEO_AllTags` |
| SEO теги отсутствуют | `TestAnalyze_SEO_NoTags` |
| HTML-сущности | `TestAnalyze_SEO_HTMLEntities` |
| Глубина обхода | `TestAnalyze_DepthZero`, `TestAnalyze_DepthOne` |
| Внешние ссылки игнорируются | `TestAnalyze_ExternalLinksIgnored` |
| Дубликаты ссылок | `TestAnalyze_DuplicateLinks` |

## CI/CD

GitHub Actions workflow в `.github/workflows/hexlet-check.yml` автоматически запускает тесты и линтер.

## Замечания и рекомендации

### Потенциальные улучшения

1. **`crawler_go_old`** — в директории есть старый файл, возможно стоит удалить
2. **Обработка `Retries`** — параметр есть в `Options`, но не используется в коде
3. **HEAD перед GET** — в `checkLink()` сначала HEAD, при ошибке GET; можно упростить
4. **`RawBody` в отчете** — передается только для внутренних нужд, но занимает память

### Сильные стороны

- ✅ Чистая архитектура с разделением на пакеты
- ✅ Полное покрытие тестами всех сценариев из ТЗ
- ✅ Правильная работа с context для отмены операций
- ✅ Детерминированные тесты без реальных HTTP-вызовов
- ✅ Поддержка rate limiting (delay + rps)
- ✅ Дедупликация URL и ограничение по домену
- ✅ Параллельная проверка битых ссылок

## Вывод

Проект **полностью соответствует требованиям ТЗ**:
- Реализованы все команды Makefile
- HTTP-логика тестируется с подменой клиента
- Поддерживается глубина обхода
- Проверяются битые ссылки
- Собираются SEO-параметры (title, description, h1)
- Используется конкурентность с горутинами
- JSON-отчет соответствует спецификации
