---
name: clean-feature-delivery
description: Implement or refactor features with clean architecture, validation, tests, and docs updates.
---

# Skill: Clean Feature Delivery

## Goal
Implement a feature using DDD + Clean Architecture with practical SOLID, without over-engineering.

## When To Use
- Add new CRUD endpoint.
- Add new use case in existing module.
- Refactor endpoint structure while preserving behavior.

## Required Steps
1. Confirm module boundary:
- Use merged-layer style (`application`, `domain`, `interfaces`, `infrastructure`) instead of per-module duplicated infrastructure.

2. Implement by layer:
- `domain`: entity/value object invariants only.
- `application`: use case orchestration + ports.
- `interfaces/http`: request/response DTO + mapper + handler + routes.
- `infrastructure/persistence`: repository implementation only.

3. Enforce standards:
- Validate request bodies with `validator/v10` tags.
- Return errors via centralized error handler contract.
- Keep timestamps in UTC.
- Keep logs structured and request-correlated.

4. Add tests:
- Unit tests for use case happy path + failure path.
- Keep tests deterministic.

5. Update docs when needed:
- Swagger annotations/routes.
- Bruno collection if endpoint contract changed.
- README/docs only when behavior or operations changed.

## Done Criteria
- `go test ./...` passes.
- No rule violation in `.cursor/rules/clean-ddd-solid.mdc`.
- Security and quality checks considered (OWASP + SonarQube aligned).
