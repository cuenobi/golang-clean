# golang-clean

A blog-ready reference project for **DDD + Clean Architecture** in Go, with practical conventions for team-scale delivery.

## Stack
- Fiber HTTP server
- Swagger/OpenAPI documentation (Swaggo)
- Kafka messaging
- Cobra CLI
- PostgreSQL + GORM
- SQL migration with golang-migrate
- Test suite with testify + mockery

## Core Conventions
- Architecture style: DDD + Clean (merged-layer style, avoid duplicated per-module infrastructure)
- SOLID: practical, not over-engineered
- Dependency rule: dependencies point inward
- UTC-only: application, database, logs, and containers
- Security baseline: OWASP Top 10 aligned
- Code quality baseline: SonarQube-style quality gate
- Error contract: stable machine-readable error codes for FE mapping

## Directory Structure
See [docs/directory-structure.md](docs/directory-structure.md) for the full up-to-date tree.

Quick placement:
- Domain model and invariants: `internal/domain`
- Use cases and ports: `internal/application`
- HTTP/messaging adapters: `internal/interfaces`
- Persistence/messaging implementation + DI: `internal/infrastructure`
- Cross-cutting concerns (`config`, `httpx`, `logger`, `validator`, `metrics`): `internal/shared`
- Generic utilities only: `pkg/utils`

## Commands
```bash
make run-api
make run-consumer
make migrate-up
make migrate-down
make test
make mockery
make swagger
make hooks-install
```

## Local Run (Docker Compose)
Configure runtime values in `.env.example`.

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

## API Documentation and Testing

### Swaggo
- Swagger UI endpoint: `GET /swagger/index.html`
- Generate docs:
```bash
make swagger
```
- Generated artifacts: `api/swagger/`

### Bruno
- Collection path: `bruno/golang-clean`
- Import this folder in Bruno and use `environments/local.bru`

## HTTP and Endpoint Conventions
- Keep handlers thin; orchestration belongs to use cases.
- Request/response DTOs stay in `internal/interfaces/http/<module>/dto.go`.
- Mapping logic stays in `mapper.go`.
- Validate request bodies with `validator/v10` tags.
- Write endpoints require JSON content type (`application/json`).
- Middleware is centralized in DI HTTP module (`internal/infrastructure/di/http_module.go`).

Current centralized controls:
- CORS (env-driven)
- Rate limit (env-driven)
- Body size limit (env-driven)
- API key auth + permission checks (when enabled)
- Request ID + structured request logging

## UTC Time Standard
UTC is enforced across the project:
- App runtime clock uses UTC
- DB timestamps use `TIMESTAMPTZ`
- DB session timezone is set to UTC
- Logger timestamps use UTC
- Container runtime uses `TZ=UTC`

## Logging (Loki Ready)
- Structured JSON logs via `zerolog` to stdout
- Request logs include `request_id`, `method`, `path`, `status`, `latency_ms`, `ip`, `user_agent`
- Levels by status:
  - `2xx/3xx`: `info`
  - `4xx`: `warn`
  - `5xx`: `error`

## Reliability and Security (Current Baseline)
- API key authentication and permission-based authorization for `/api/v1/*`
- Rate limiting middleware
- Operational endpoints:
  - `GET /healthz`
  - `GET /readyz`
  - `GET /metrics`
- Outbox pattern for `order.created.v1` event delivery
- Consumer dispatcher with retry/backoff, publish timeout, and circuit breaker
- Idempotent order creation via `Idempotency-Key`

## Error Handling Contract (FE Mapping)
All API errors use one stable shape:

```json
{
  "code": 1001,
  "type": "validation_error",
  "message": "email failed on 'email'",
  "data": {
    "violations": [
      { "field": "email", "rule": "email" }
    ]
  },
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
- `1413` = `payload_too_large`
- `1415` = `unsupported_media_type`
- `1422` = `invalid_state`
- `1429` = `rate_limited`

Context-specific codes are defined in `internal/shared/kernel/error_codes_context.go`.

## Environment Variables
Full template: `.env.example`

- App: `APP_NAME`, `APP_ENV`, `APP_TIMEZONE`, `LOG_LEVEL`
- HTTP: `HTTP_ADDRESS`, `HTTP_READ_TIMEOUT_SEC`, `HTTP_WRITE_TIMEOUT_SEC`, `HTTP_BODY_LIMIT_MB`, `READINESS_DB_TIMEOUT_MS`
- Security: `AUTH_ENABLED`, `API_KEY`, `RATE_LIMIT_ENABLED`, `RATE_LIMIT_MAX`, `RATE_LIMIT_WINDOW_SEC`
- CORS: `CORS_ALLOWED_ORIGINS`, `CORS_ALLOWED_METHODS`, `CORS_ALLOWED_HEADERS`, `CORS_EXPOSE_HEADERS`, `CORS_ALLOW_CREDENTIALS`, `CORS_MAX_AGE_SEC`
- PostgreSQL: `POSTGRES_HOST`, `POSTGRES_PORT`, `POSTGRES_USER`, `POSTGRES_PASS`, `POSTGRES_DB`, `POSTGRES_SSL_MODE`, `POSTGRES_TIMEZONE`
- Kafka: `KAFKA_BROKER`, `KAFKA_PUBLISH_TIMEOUT_MS`
- Outbox: `OUTBOX_POLL_INTERVAL_MS`, `OUTBOX_BATCH_SIZE`, `OUTBOX_MAX_RETRIES`, `OUTBOX_RETRY_BACKOFF_MS`, `OUTBOX_PROCESSING_TIMEOUT_MS`
- Circuit breaker: `CIRCUIT_BREAKER_FAILURES`, `CIRCUIT_BREAKER_OPEN_MS`

## Migration Policy
- Use `golang-migrate` SQL files in `migrations/`
- Do not use GORM AutoMigrate in production
- Current baseline migration is consolidated in `000001_init.(up|down).sql`

## Team Commit Standard
- Pattern: `<type>(<scope>): <subject>`
- Rule: `.cursor/rules/commit-message-standard.mdc`
- Template: `.gitmessage`
- Hook: `.githooks/commit-msg`
- Install:
```bash
make hooks-install
```

## Cursor Rules and Skills

### Rules
- `.cursor/rules/clean-ddd-solid.mdc`
- `.cursor/rules/security-owasp-top10.mdc`
- `.cursor/rules/code-quality-sonarqube.mdc`
- `.cursor/rules/http-api-convention.mdc`
- `.cursor/rules/utc-time-standard.mdc`
- `.cursor/rules/commit-message-standard.mdc`

### Skills
- `.cursor/skills/clean-feature-delivery.mdc`
- `.cursor/skills/owasp-security-review.mdc`
- `.cursor/skills/sonarqube-quality-gate.mdc`
- `.cursor/skills/release-readiness-check.mdc`

## Related Docs
- [docs/architecture.md](docs/architecture.md)
- [docs/directory-structure.md](docs/directory-structure.md)
- [.cursor/skills/README.md](.cursor/skills/README.md)
