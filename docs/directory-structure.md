# Directory Structure (Current)

```text
.
├── .cursor/
│   ├── rules/                         # Cursor governance rules (architecture, security, quality, UTC, commit)
│   └── skills/                        # Reusable prompt skills for team workflows
├── .env.example                       # Environment variable template
├── .githooks/                         # Git hooks (commit-msg validation)
├── .gitmessage                        # Git commit template
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
│   ├── application/                   # Application layer
│   │   ├── dto/                       # DTOs grouped by module (order, user)
│   │   ├── port/
│   │   │   ├── in/                    # Input ports (use case contracts)
│   │   │   └── out/                   # Output ports (repo/contracts)
│   │   └── usecase/                   # Use cases grouped by module
│   │       ├── order/                 # create/get/list/update/delete split files
│   │       └── user/                  # create/get/list/update/delete split files
│   │
│   ├── domain/                        # Entities, value objects, domain events
│   │   ├── entity/
│   │   ├── valueobject/
│   │   └── event/
│   │
│   ├── interfaces/                    # Interface adapters
│   │   ├── http/
│   │   │   ├── order/                 # handler + dto + mapper + routes
│   │   │   ├── user/                  # handler + dto + mapper + routes
│   │   │   └── system/                # health/readiness/metrics endpoints
│   │   └── messaging/                 # Messaging consumer adapters
│   │
│   ├── infrastructure/                # Technical implementations
│   │   ├── di/                        # Composition root modules (container/http/messaging/module)
│   │   ├── messaging/                 # Kafka publisher adapters
│   │   └── persistence/               # GORM repository adapters + outbox storage
│   │
│   └── shared/                        # Cross-cutting components
│       ├── config/                    # env config + DSN/migration URL
│       ├── httpx/                     # error handler, auth, CORS, rate limit, request middleware
│       ├── kernel/                    # app errors, error codes, contracts/providers
│       ├── logger/                    # zerolog structured logging
│       ├── metrics/                   # prometheus middleware
│       ├── persistence/               # gorm setup + tx manager/context
│       ├── resilience/                # circuit breaker
│       ├── validator/                 # validator v10 + field violations
│       └── kafkax/                    # kafka config helpers
│
├── api/
│   ├── openapi/                       # Static OpenAPI YAML
│   └── swagger/                       # Swaggo generated docs
│
├── migrations/                        # golang-migrate SQL files (consolidated baseline: 000001)
├── docs/                              # Architecture and structure docs
├── scripts/                           # Local setup scripts (hooks install)
├── tests/                             # Integration/contract placeholders
├── pkg/utils/                         # Generic utility helpers
└── tools/                             # Tool dependencies (mockery, swag)
```

## Placement Guide

```text
Business entity/value object   -> internal/domain/
Use case orchestration         -> internal/application/usecase/
HTTP endpoint adapter          -> internal/interfaces/http/<module>/
Persistence adapter            -> internal/infrastructure/persistence/
Cross-cutting concerns         -> internal/shared/
Generic helper (non-business)  -> pkg/utils/
```

## Use Case File Convention

```text
internal/application/usecase/order/
  usecase.go
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

Same convention applies to `internal/application/usecase/user/`.
