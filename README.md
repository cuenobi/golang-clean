# golang-clean

A blog-ready reference project for **DDD + Clean Architecture** in Go.

## Stack
- Fiber HTTP server
- Swagger/OpenAPI documentation
- Kafka messaging
- Cobra CLI
- PostgreSQL + GORM
- SQL migration with golang-migrate
- Test suite with testify + mockery

## Project Principles
- Dependency rule: code points inward.
- Domain must not import frameworks.
- Application owns use cases and ports.
- Infrastructure only implements ports.
- Adapters map transport models to application DTOs.
- Keep `pkg/utils` framework-agnostic and business-agnostic.

## Commands
```bash
make run-api
make run-consumer
make migrate-up
make migrate-down
make test
```

## Logging (Loki Ready)
- Structured JSON logs via `zerolog` to `stdout` (recommended pattern for Grafana Loki/Promtail).
- HTTP request logs include `request_id`, `method`, `path`, `status`, `latency_ms`, `ip`, `user_agent`.
- Log levels by HTTP status:
  - `2xx/3xx`: `info`
  - `4xx`: `warn`
  - `5xx`: `error`

Environment variables:
- `APP_NAME` (default: `golang-clean`)
- `APP_ENV` (default: `dev`)
- `LOG_LEVEL` (default: `info`, supported: `debug`, `info`, `warn`, `error`)

## Architecture Docs
See [docs/architecture.md](docs/architecture.md), [docs/directory-structure.md](docs/directory-structure.md), and [.cursor/rules/clean-ddd-solid.mdc](.cursor/rules/clean-ddd-solid.mdc).
