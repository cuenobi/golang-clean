# ADR-001: Merged Layer DDD + Clean Structure

- Status: Accepted
- Date: 2026-04-18

## Context
The project started with context-silo style directories that duplicated infrastructure per domain and caused overhead for a single service scope.

## Decision
Use a merged-layer structure:
- `internal/domain`
- `internal/application`
- `internal/interfaces`
- `internal/infrastructure`
- `internal/shared`

Keep business domains (`user`, `order`) separated by modules inside layers, not by duplicated full stacks.

## Consequences
### Positive
- Lower operational and maintenance overhead.
- Clear Clean Architecture dependency direction.
- Easier onboarding and blog-level explainability.

### Negative
- Requires strict conventions to prevent layer leakage.
- Larger shared layers can become dumping grounds without reviews.

## Guardrails
- Domain must not import framework libraries.
- Handlers stay thin; use cases own orchestration.
- Shared package must remain cross-cutting only.
