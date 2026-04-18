# DDD + Clean Architecture Guide

## Layer Boundaries
1. Domain
- Entities, value objects, domain events, invariants.
- No dependency on Fiber, Kafka, GORM, Cobra, Swagger.

2. Application
- Use cases and orchestration.
- Defines inbound/outbound ports.
- Transaction and idempotency boundary.

3. Interfaces (Adapters)
- HTTP handlers (Fiber), Kafka consumers.
- Mapping request/response payloads to DTOs.

4. Infrastructure
- Postgres/GORM repositories.
- Kafka producers/consumers.
- CLI and bootstrap wiring.

## SOLID (practical, not over-engineered)
- Single Responsibility: handlers map and delegate; use cases own business orchestration.
- Open/Closed: new adapter can be added by implementing ports.
- Liskov Substitution: interfaces must preserve behavior contracts.
- Interface Segregation: small focused interfaces (`OrderReader`, `OrderWriter`).
- Dependency Inversion: application depends on abstractions, infra depends on application ports.

## Migration Policy
- Use `golang-migrate` SQL files under `migrations/`.
- No GORM AutoMigrate in production.

## Microservice Readiness
- Each bounded context stays under `internal/<context>`.
- Context can be extracted later with minimal rewrite.

## Lean Refactor Guideline
- Keep domain and usecase separated by bounded context (`internal/order`, `internal/user`).
- Share technical primitives in `internal/shared` when they are truly cross-context:
  - `Clock`, `IDGenerator`, transaction manager, DB bootstrap, logger.
- Do not duplicate infrastructure boilerplate per context unless behavior is different.
- Create context-specific infrastructure only when business or integration needs diverge.
