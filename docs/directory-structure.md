# Directory Structure (Current)

```text
.
├── .cursor/rules/                     # Cursor rules (team conventions)
├── .env.example                       # Environment variable template
├── docker-compose.yml                 # Local stack (postgres, kafka, api, consumer, migrate)
├── Dockerfile                         # Service container build
├── bruno/golang-clean/                # Bruno API collection + environments
│
├── cmd/                               # Cobra entry commands
│   └── app/main.go                    # Application entrypoint
│
├── internal/
│   ├── bootstrap/                     # App bootstrap (config + db + module wiring)
│   │
│   ├── application/                   # Application layer (order + user together)
│   │   ├── dto/                       # DTOs grouped by module
│   │   │   ├── order/
│   │   │   └── user/
│   │   ├── port/
│   │   │   ├── in/                    # Input ports (use case contracts)
│   │   │   └── out/                   # Output ports (repo/contracts)
│   │   └── usecase/                   # Use cases grouped by module
│   │       ├── order/
│   │       └── user/
│   │
│   ├── domain/                        # Domain layer (entities/value objects/events)
│   │   ├── entity/
│   │   ├── valueobject/
│   │   └── event/
│   │
│   ├── interfaces/                    # Interface adapters
│   │   ├── http/
│   │   │   ├── order/                 # Order HTTP handlers
│   │   │   ├── system/                # Health/readiness/metrics endpoints
│   │   │   └── user/                  # User HTTP handlers
│   │   └── messaging/                 # Messaging consumer adapters
│   │
│   ├── infrastructure/                # Technical implementations
│   │   ├── di/                        # Modular composition root (container + modules)
│   │   ├── messaging/                 # Kafka publisher adapters
│   │   └── persistence/               # GORM repository adapters + outbox storage
│   │
│   └── shared/                        # Shared cross-cutting components
│       ├── config/
│       ├── httpx/
│       ├── kafkax/
│       ├── kernel/
│       ├── logger/
│       ├── metrics/
│       ├── persistence/
│       ├── resilience/
│       └── validator/
│
├── api/
│   ├── openapi/                       # Static OpenAPI YAML
│   └── swagger/                       # Swaggo generated docs
│
├── migrations/                        # SQL migrations
├── docs/                              # Project architecture docs
├── scripts/                           # Utility scripts for local development
├── tests/                             # Integration/contract test placeholders
├── pkg/utils/                         # Generic utility helpers
└── tools/                             # Tool dependencies (mockery, swag)
```

## Quick Placement Guide

```text
Business entity/value object   -> internal/domain/
Use case logic                 -> internal/application/usecase/
HTTP endpoint                  -> internal/interfaces/http/<resource>/
DB repository implementation   -> internal/infrastructure/persistence/
Cross-cutting infra/utilities  -> internal/shared/
```

## Usecase File Convention

```text
internal/application/usecase/order/
  usecase.go         # type + constructor + dependencies
  create.go
  get.go
  list.go
  update.go
  delete.go
  mapper.go
  create_test.go
  get_test.go
  list_test.go
  update_test.go
  delete_test.go
  test_helpers_test.go
```
