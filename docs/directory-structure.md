# Directory Structure (Current)

```text
.
├── .cursor/
│   ├── commands/                      # Cursor command playbooks
│   ├── rules/                         # Governance rules (architecture, security, quality, API, UTC, commit)
│   └── skills/                        # Reusable team skills (readiness, tests, security, release)
├── .env.example                       # Environment variable template
├── .githooks/
│   └── commit-msg                     # Commit message validation hook
├── .gitmessage                        # Git commit template
├── Dockerfile                         # Service container build
├── docker-compose.yml                 # Local stack (postgres, kafka, api, consumer, migrate)
├── Jenkinsfile                        # CI pipeline definition
├── Makefile                           # Common local/CI commands
├── go.mod                             # Go module definition
├── go.sum                             # Go dependency checksums
├── bruno/golang-clean/                # Bruno API collection + local environment
│   ├── Audit Logs/
│   ├── Orders/
│   ├── Users/
│   ├── System/
│   └── environments/
│
├── cmd/                               # Cobra CLI commands
│   ├── app/main.go                    # Application entrypoint
│   ├── root.go                        # Root command wiring
│   ├── api.go                         # Start HTTP API command
│   ├── consumer.go                    # Start messaging consumer command
│   ├── migrate.go                     # Run DB migrations command
│   └── api_docs.go                    # API docs command
│
├── internal/
│   ├── bootstrap/                     # App bootstrap (config + db + module wiring)
│   ├── application/                   # Application layer
│   │   ├── dto/                       # DTOs grouped by module (order, user, auditlog)
│   │   ├── port/
│   │   │   ├── in/                    # Input ports (use case contracts)
│   │   │   └── out/                   # Output ports (repo/publisher contracts)
│   │   └── usecase/                   # Use cases grouped by module (order, user, auditlog)
│   ├── domain/                        # Entities, value objects, domain events
│   │   ├── entity/
│   │   ├── valueobject/
│   │   └── event/
│   ├── interfaces/                    # Interface adapters
│   │   ├── http/
│   │   │   ├── order/                 # handler + dto + mapper + routes
│   │   │   ├── user/                  # handler + dto + mapper + routes
│   │   │   ├── system/                # health/readiness endpoints
│   │   │   └── auditlog/              # handler + dto + mapper + routes
│   │   └── messaging/                 # Messaging consumer adapters
│   ├── infrastructure/                # Technical implementations
│   │   ├── di/                        # Composition root modules (container/http/messaging/module)
│   │   ├── messaging/                 # Kafka publisher adapters
│   │   └── persistence/               # GORM models/repositories + outbox storage
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
├── migrations/                        # golang-migrate SQL files
├── docs/                              # Architecture and structure docs
├── scripts/                           # Local setup scripts
├── tests/
│   ├── contract/                      # Contract test placeholder/docs
│   └── integration/                   # Integration test placeholder/docs
├── pkg/utils/                         # Generic utility helpers
└── tools/                             # Tool dependencies (mockery, swag)
```

## Placement Guide

```text
Business entity/value object   -> internal/domain/
Use case orchestration         -> internal/application/usecase/<module>/
HTTP endpoint adapter          -> internal/interfaces/http/<module>/
Messaging adapter              -> internal/interfaces/messaging/ and internal/infrastructure/messaging/
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

Same convention applies to `internal/application/usecase/user/` and `internal/application/usecase/auditlog/`.
