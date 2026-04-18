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

Environment variables:
- `APP_NAME` (default: `golang-clean`)
- `APP_ENV` (default: `dev`)
- `LOG_LEVEL` (default: `info`, supported: `debug`, `info`, `warn`, `error`)
- `AUTH_ENABLED` (default: `false`)
- `API_KEY` (required when `AUTH_ENABLED=true`)
- `RATE_LIMIT_ENABLED` (default: `true`)
- `RATE_LIMIT_MAX` (default: `120`)
- `RATE_LIMIT_WINDOW_SEC` (default: `60`)
- `POSTGRES_SSL_MODE` (default: `disable`)
- `OUTBOX_POLL_INTERVAL_MS` (default: `1000`)
- `OUTBOX_BATCH_SIZE` (default: `50`)
- `OUTBOX_MAX_RETRIES` (default: `5`)
- `OUTBOX_RETRY_BACKOFF_MS` (default: `500`)
- `OUTBOX_PROCESSING_TIMEOUT_MS` (default: `15000`)
- `KAFKA_PUBLISH_TIMEOUT_MS` (default: `3000`)
- `CIRCUIT_BREAKER_FAILURES` (default: `5`)
- `CIRCUIT_BREAKER_OPEN_MS` (default: `30000`)

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
- [docs/phase2/README.md](docs/phase2/README.md)
- [docs/phase2/nfr-sla.md](docs/phase2/nfr-sla.md)
- [docs/phase2/ownership-oncall.md](docs/phase2/ownership-oncall.md)
- [docs/phase2/risk-tradeoff-matrix.md](docs/phase2/risk-tradeoff-matrix.md)
- [docs/phase2/error-code-catalog.md](docs/phase2/error-code-catalog.md)
- [.cursor/rules/clean-ddd-solid.mdc](.cursor/rules/clean-ddd-solid.mdc)
