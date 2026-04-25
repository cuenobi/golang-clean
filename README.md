# golang-clean

DDD + Clean Architecture example in Go, designed for AI-assisted development (Cursor) with practical team conventions.

## What You Get
- Fiber HTTP API
- Kafka consumer + outbox pattern
- PostgreSQL + GORM
- `golang-migrate` SQL migrations
- Swagger docs (Swaggo)
- Bruno collection
- Audit log endpoints (`/api/v1/audit-logs/system`, `/api/v1/audit-logs/organization`)
- Structured logging (Loki-friendly)
- UTC-only time policy
- OWASP + SonarQube aligned project rules

## Prerequisites
- Go `1.25+`
- Docker + Docker Compose
- Make

## Quick Start (Recommended: Docker)

1. Clone repository
```bash
git clone https://github.com/cuenobi/golang-clean.git
cd golang-clean
```

2. Review environment template
```bash
cat .env.example
```

3. Run DB migration (starts PostgreSQL automatically)
```bash
docker compose --profile tools up --build migrate
```

4. Start API + Consumer
```bash
docker compose up --build api consumer
```

5. Verify services
```bash
curl -i http://localhost:8080/healthz
curl -i http://localhost:8080/readyz
curl -i http://localhost:8080/metrics
```

6. Open Swagger UI
- `http://localhost:8080/swagger/index.html`

## Local Run (Go Process)
Use this mode if you want to run `go run` directly.

1. Start infrastructure only
```bash
docker compose up -d postgres kafka
```

2. Export env vars for local process
```bash
set -a
source .env.example
set +a

# local overrides (containers are exposed on localhost)
export POSTGRES_HOST=localhost
export KAFKA_BROKER=localhost:9092
```

3. Run migration
```bash
make migrate-up
```

4. Run API and consumer (separate terminals)
```bash
make run-api
make run-consumer
```

## Common Commands
```bash
make run-api
make run-consumer
make migrate-up
make migrate-down
make test
make mockery
make swagger
make hooks-install
make fmt
```

## API Auth Notes
- Default in `.env.example`: `AUTH_ENABLED=false`
- If you enable auth (`AUTH_ENABLED=true`):
  - send `X-API-Key`
  - send `X-Permissions` per endpoint (for example `users:read`, `orders:write`, `audit_logs:read`)

## API Error Contract
All API errors use this shape:

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
  "request_id": "..."
}
```

Common codes:
- `1000` internal
- `1001` validation
- `1002` bad request
- `1401` unauthorized
- `1403` forbidden
- `1404` not found
- `1409` conflict
- `1413` payload too large
- `1415` unsupported media type
- `1422` invalid state
- `1429` rate limited

## Documentation
- Architecture: [docs/architecture.md](docs/architecture.md)
- Directory structure: [docs/directory-structure.md](docs/directory-structure.md)

## Bruno
- Collection: `bruno/golang-clean`
- Environment: `bruno/golang-clean/environments/local.bru`

## Team Conventions (Cursor)
### Rules
- `.cursor/rules/clean-ddd-solid.mdc`
- `.cursor/rules/http-api-convention.mdc`
- `.cursor/rules/utc-time-standard.mdc`
- `.cursor/rules/security-owasp-top10.mdc`
- `.cursor/rules/code-quality-sonarqube.mdc`
- `.cursor/rules/commit-message-standard.mdc`
- `.cursor/rules/endpoint-delivery-artifacts.mdc`

### Skills
- `.cursor/skills/clean-feature-delivery/SKILL.md`
- `.cursor/skills/owasp-security-review/SKILL.md`
- `.cursor/skills/sonarqube-quality-gate/SKILL.md`
- `.cursor/skills/release-readiness-check/SKILL.md`
- `.cursor/skills/release-skills/SKILL.md`
- `.cursor/skills/unit-test-mockery-suite/SKILL.md`

### Commands
- `.cursor/commands/implement-ticket.md`
- `.cursor/commands/generate-unit-test.md`

## Git Hooks / Commit Template
Install project commit hooks and commit message template:

```bash
make hooks-install
```

## Troubleshooting
- `connection refused` to Postgres/Kafka:
  - ensure `docker compose up -d postgres kafka` is running
- migration fails:
  - verify `POSTGRES_*` env values
- API returns `unsupported_media_type`:
  - send `Content-Type: application/json` on write endpoints
- API returns `payload_too_large`:
  - increase `HTTP_BODY_LIMIT_MB` if needed
