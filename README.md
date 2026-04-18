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

## Architecture Docs
See [docs/architecture.md](docs/architecture.md), [docs/directory-structure.md](docs/directory-structure.md), and [.cursor/rules/clean-ddd-solid.mdc](.cursor/rules/clean-ddd-solid.mdc).
