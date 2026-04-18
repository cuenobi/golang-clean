# Risk and Trade-off Matrix

| Area | Risk | Impact | Mitigation | Trade-off |
|---|---|---|---|---|
| DB + Event consistency | Event lost/duplicated under partial failure | High | Outbox pattern + idempotency key + retries | More operational complexity |
| Auth disabled in non-dev | Public API exposure | High | Require `AUTH_ENABLED=true` in SIT/UAT/PROD and enforce env policy | Harder local parity if env misconfigured |
| Rate limiting too strict | Legit traffic throttled | Medium | Tune by env and monitor `429` ratio | Higher abuse risk if too loose |
| Circuit breaker too sensitive | Temporary stop of event publishing | Medium | Tune failure threshold/open duration | Potential delay in message propagation |
| Schema migration errors | Deployment rollback difficulty | High | Forward-only checks + tested down migrations | Slower release process |
| Shared layer bloat | Architecture erosion | Medium | Strict code review checklist and ADR discipline | More review overhead |
| Logging volume growth | Cost increase in Loki stack | Medium | Sampling strategy, level tuning by env | Lower granularity for debugging |
| Missing integration tests | Regression escapes to runtime | Medium | Add DB/Kafka integration and contract tests in CI | Increased pipeline time |

## Accepted Trade-offs
1. Eventual consistency for domain events is accepted to improve reliability.
2. Slightly stricter route middleware complexity is accepted for production safety.
3. Extra governance docs are accepted to keep architectural decisions explicit.
