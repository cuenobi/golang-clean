# Directory Structure (Merged Layer Style)

```text
.
├── cmd/                               # Cobra entry commands
│   └── app/main.go                    # Application entrypoint
│
├── internal/
│   ├── bootstrap/                     # App bootstrap (config + db + module wiring)
│   │
│   ├── application/                   # Application layer (order + user together)
│   │   ├── dto/                       # Request/response DTOs
│   │   ├── port/
│   │   │   ├── in/                    # Input ports (use case contracts)
│   │   │   └── out/                   # Output ports (repo/contracts)
│   │   └── usecase/                   # Use case implementations
│   │
│   ├── domain/                        # Domain layer (entities/value objects/events)
│   │   ├── entity/
│   │   ├── valueobject/
│   │   └── event/
│   │
│   ├── interfaces/                    # Interface adapters
│   │   ├── http/
│   │   │   ├── order/                 # Order HTTP handlers
│   │   │   └── user/                  # User HTTP handlers
│   │   └── messaging/                 # Messaging consumer adapters
│   │
│   ├── infrastructure/                # Technical implementations
│   │   ├── di/                        # Module composition root
│   │   ├── messaging/                 # Kafka publisher adapters
│   │   └── persistence/               # GORM repository adapters
│   │
│   └── shared/                        # Shared cross-cutting components
│       ├── config/
│       ├── httpx/
│       ├── kafkax/
│       ├── kernel/
│       ├── logger/
│       ├── persistence/
│       └── validator/
│
├── api/                               # OpenAPI/Swagger docs
├── migrations/                        # SQL migrations
├── docs/                              # Project docs
├── tests/                             # Integration/contract test placeholders
├── pkg/utils/                         # Generic utility helpers
└── tools/                             # Tool dependencies (mockery)
```

## Quick Placement Guide

```text
Business entity/value object   -> internal/domain/
Use case logic                 -> internal/application/usecase/
HTTP endpoint                  -> internal/interfaces/http/<resource>/
DB repository implementation   -> internal/infrastructure/persistence/
Cross-cutting infra/utilities  -> internal/shared/
```
