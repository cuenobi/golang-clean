# ADR-002: Outbox + Idempotency for Order Created Event

- Status: Accepted
- Date: 2026-04-18

## Context
Publishing Kafka messages directly in use case flow can lead to inconsistency between DB state and message delivery under partial failure.

## Decision
1. Persist order + outbox record in the same DB transaction.
2. Dispatch outbox asynchronously from consumer loop.
3. Apply retry/backoff, publish timeout, and circuit breaker in dispatcher.
4. Support idempotent `CreateOrder` by `Idempotency-Key` header.

## Consequences
### Positive
- Stronger consistency guarantees at service boundary.
- Retries become explicit and observable.
- Reduced duplicate order creation from client retries.

### Negative
- Additional table and operational complexity.
- Event delivery is asynchronous (eventual consistency).

## Operational Notes
- Failed events are marked `DEAD` after max retries.
- `OUTBOX_*` and `CIRCUIT_BREAKER_*` env values tune behavior.
