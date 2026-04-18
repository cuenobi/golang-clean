# Phase 2 Governance Pack

This folder defines architecture governance for production and team operations.

## Contents
- `adr/`: Architecture Decision Records (ADRs)
- `nfr-sla.md`: Non-functional requirements and SLO/SLA targets
- `ownership-oncall.md`: Service ownership and support model
- `risk-tradeoff-matrix.md`: Technical/business risks with documented trade-offs
- `error-code-catalog.md`: Stable FE-facing error code contract by context

## How To Use
1. Keep these docs versioned with code changes.
2. Any architectural change must include:
   - ADR update (new or amended)
   - risk/trade-off update
   - error code update (if API error behavior changes)
3. Keep all numeric error codes backward compatible once public.
