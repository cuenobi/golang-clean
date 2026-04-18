# DDD + Clean Architecture Guide

## Current Project Style
This project now uses a **merged layer style** to reduce over-engineering:
- `order` and `user` live in the same layer folders (`application`, `domain`, `interfaces`, `infrastructure`)
- shared technical concerns are centralized in `internal/shared`

## Layer Boundaries
1. Domain
- Entities, value objects, domain events, invariants.
- No dependency on Fiber, Kafka, GORM, Cobra, Swagger.

2. Application
- Use cases and orchestration.
- Defines inbound/outbound ports.

3. Interfaces (Adapters)
- HTTP handlers (Fiber), Kafka consumers.
- Mapping transport payloads to application DTOs.

4. Infrastructure
- Postgres/GORM repository implementations.
- Kafka producers/consumers.
- Composition root and wiring.

## SOLID (Practical)
- SRP: handlers are thin; usecases own business orchestration.
- OCP: add adapters by implementing ports/contracts.
- ISP: keep interfaces small and role-focused.
- DIP: application depends on abstractions; infrastructure depends on application/domain contracts.

## Migration Policy
- Use `golang-migrate` SQL files under `migrations/`.
- No GORM AutoMigrate in production.

## Logging Standard
- Use structured JSON logs to `stdout` for container-first deployments.
- Set log level via environment variable `LOG_LEVEL`.
- Include request correlation fields (`request_id`, method/path/status, latency).
- This format is designed for log shipping to Grafana Loki via Promtail/Alloy.

## Why This Structure
- Avoids duplicating infrastructure boilerplate per bounded context.
- Keeps clean-layer boundaries while staying practical for a small/medium service.
- Easier to publish and maintain without deep directory nesting.
