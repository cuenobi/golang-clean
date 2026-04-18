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
make swagger
```

## Docker Compose (Local Run)
Configure runtime values in `config.example` (used directly by `docker-compose.yml`).

1. Start dependencies + run migrations:
```bash
docker compose --profile tools up --build migrate
```
2. Start API + Consumer:
```bash
docker compose up --build api consumer
```
3. Optional full stack:
```bash
docker compose up --build
```

## Swaggo
- Swagger UI endpoint: `GET /swagger/index.html`
- Generate docs:
```bash
make swagger
```
- Generated artifacts live in `api/swagger/`.

## Bruno
- Bruno collection path: `bruno/golang-clean`
- Import this folder in Bruno and use `environments/local.bru`.

## Phase 1 Hardening
- Authentication with API Key (`X-API-Key`) and permission-based authorization (`X-Permissions`).
- Rate limiting middleware with configurable window and quota.
- Operational endpoints:
  - `GET /healthz` (liveness)
  - `GET /readyz` (readiness with DB ping)
  - `GET /metrics` (Prometheus)
- Outbox pattern for order-created events (DB write + event enqueue in same transaction).
- Idempotent order creation with `Idempotency-Key` header.
- Consumer retry/backoff + publish timeout + circuit breaker + dead status.

## Logging (Loki Ready)
- Structured JSON logs via `zerolog` to `stdout` (recommended pattern for Grafana Loki/Promtail).
- HTTP request logs include `request_id`, `method`, `path`, `status`, `latency_ms`, `ip`, `user_agent`.
- Log levels by HTTP status:
  - `2xx/3xx`: `info`
  - `4xx`: `warn`
  - `5xx`: `error`

Environment variables (full template: `config.example`):
- App: `APP_NAME`, `APP_ENV`, `LOG_LEVEL`
- HTTP: `HTTP_ADDRESS`, `HTTP_READ_TIMEOUT_SEC`, `HTTP_WRITE_TIMEOUT_SEC`, `READINESS_DB_TIMEOUT_MS`
- Security: `AUTH_ENABLED`, `API_KEY`, `RATE_LIMIT_ENABLED`, `RATE_LIMIT_MAX`, `RATE_LIMIT_WINDOW_SEC`
- PostgreSQL: `POSTGRES_HOST`, `POSTGRES_PORT`, `POSTGRES_USER`, `POSTGRES_PASS`, `POSTGRES_DB`, `POSTGRES_SSL_MODE`
- Kafka: `KAFKA_BROKER`, `KAFKA_PUBLISH_TIMEOUT_MS`
- Outbox: `OUTBOX_POLL_INTERVAL_MS`, `OUTBOX_BATCH_SIZE`, `OUTBOX_MAX_RETRIES`, `OUTBOX_RETRY_BACKOFF_MS`, `OUTBOX_PROCESSING_TIMEOUT_MS`
- Circuit breaker: `CIRCUIT_BREAKER_FAILURES`, `CIRCUIT_BREAKER_OPEN_MS`

## Error Handling Contract (FE Mapping)
All API errors are returned in a stable shape:

```json
{
  "code": 1001,
  "type": "validation_error",
  "message": "email failed on 'email'",
  "request_id": "8a1f..."
}
```

Common error codes:
- `1000` = `internal_error`
- `1001` = `validation_error`
- `1002` = `bad_request`
- `1401` = `unauthorized`
- `1403` = `forbidden`
- `1404` = `not_found`
- `1409` = `conflict`
- `1422` = `invalid_state`
- `1429` = `rate_limited`

## Architecture Docs
See:
- [docs/architecture.md](docs/architecture.md)
- [docs/directory-structure.md](docs/directory-structure.md)
- [.cursor/rules/clean-ddd-solid.mdc](.cursor/rules/clean-ddd-solid.mdc)
