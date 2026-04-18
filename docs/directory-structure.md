# Directory Structure (Tree Style)

```text
.
├── cmd/                            # Application commands (Cobra)
│   └── app/                        # Binary entrypoint
│       └── main.go                 # Starts root command
│
├── internal/                       # Private application code
│   ├── bootstrap/                  # App bootstrap wiring (config, db, module)
│   │
│   ├── shared/                     # Shared technical components (no business logic)
│   │   ├── config/                 # Environment/config loader
│   │   ├── logger/                 # Logging abstraction
│   │   ├── httpx/                  # HTTP shared helpers (error mapping, middleware utils)
│   │   ├── kafkax/                 # Kafka shared config/helpers
│   │   ├── persistence/            # Shared DB helpers (gorm init, tx context)
│   │   ├── kernel/                 # Shared primitives/errors
│   │
│   └── order/                      # Bounded context: Order
│       ├── domain/                 # Domain layer (business rules only)
│       │   ├── entity/             # Domain entities
│       │   ├── valueobject/        # Value objects
│       │   ├── event/              # Domain events
│       │
│       ├── application/            # Application layer (use case orchestration)
│       │   ├── usecase/            # Use case implementations
│       │   ├── port/
│       │   │   ├── in/             # Input ports (called by adapters)
│       │   │   └── out/            # Output ports (implemented by infra)
│       │   ├── dto/                # Application DTOs
│       │
│       ├── interfaces/             # Interface adapters
│       │   ├── http/               # Fiber handlers + request/response mappers
│       │   └── messaging/          # Kafka consumer adapters + mappers
│       │
│       └── infrastructure/         # Technical implementations
│           ├── persistence/        # GORM repository implementations
│           ├── messaging/          # Kafka publisher/producer implementations
│           └── di/                 # Dependency wiring for module
│
│   └── user/                       # Bounded context: User (same layered pattern)
│       ├── domain/
│       ├── application/
│       ├── interfaces/
│       └── infrastructure/
│
├── api/                            # API documentation assets
│   ├── openapi/                    # OpenAPI source file(s)
│   └── swagger/                    # Generated swagger artifacts
│
├── migrations/                     # SQL migrations (golang-migrate)
│
├── tests/                          # Higher-level tests
│   ├── integration/                # Integration tests
│   └── contract/                   # API/event contract tests
│
├── pkg/
│   └── utils/                      # Generic reusable utils (framework/business agnostic)
│
├── docs/                           # Project documentation
├── scripts/                        # Dev scripts
└── tools/                          # Tooling dependencies (e.g. mockery)
```

## Quick Placement Guide

```text
New business rule?              -> internal/<context>/domain/
New use case?                   -> internal/<context>/application/usecase/
New repository interface?       -> internal/<context>/application/port/out/
New DB/Kafka implementation?    -> internal/<context>/infrastructure/
New HTTP endpoint?              -> internal/<context>/interfaces/http/
Generic helper only?            -> pkg/utils/
```
