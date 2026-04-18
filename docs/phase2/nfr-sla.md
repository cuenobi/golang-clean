# NFR and SLA Targets

## Scope
Applies to HTTP APIs and asynchronous outbox event delivery for this service.

## Availability
- API availability target (monthly): `99.9%`
- Readiness endpoint must fail fast when DB dependency is unavailable.

## Latency
- `GET` APIs p95: `< 150 ms`
- `POST/PUT/DELETE` APIs p95: `< 300 ms`
- Readiness probe response: `< 100 ms` in healthy state

## Error Budget
- Monthly error budget for 99.9%: ~43 minutes downtime.
- If burned >50% mid-cycle, freeze non-critical changes.

## Throughput
- Baseline target: `>= 100 RPS` per instance for CRUD APIs under nominal payload.
- Rate limiter defaults:
  - `RATE_LIMIT_MAX=120`
  - `RATE_LIMIT_WINDOW_SEC=60`

## Consistency Model
- CRUD read/write in primary DB: strong consistency at transaction boundary.
- Event publishing via outbox: eventual consistency.

## Security NFR
- When `AUTH_ENABLED=true`, all `/api/v1/*` require `X-API-Key`.
- Authorization by `X-Permissions`.
- Input validation must be enforced at transport layer.
- TLS for DB should be enabled outside local dev (`POSTGRES_SSL_MODE=require` or stronger).

## Observability NFR
- Structured logs to stdout JSON.
- Every request carries `request_id`.
- Expose `/healthz`, `/readyz`, `/metrics`.

## Recovery Targets
- RTO target: `<= 30 minutes`
- RPO target: `<= 5 minutes` for data stores with managed backup strategy
