# Ownership and On-call Model

## Service Ownership
- Team: `OMS Core Backend`
- Service: `golang-clean` (reference + production-template baseline)
- Primary language: Go

## Responsibilities
- Maintain API contract compatibility.
- Keep migration and rollback paths safe.
- Maintain outbox delivery health.
- Keep operational docs current.

## On-call Rotation
- Primary on-call: backend engineer of the week.
- Secondary on-call: backup backend engineer.
- Escalation:
  1. Primary on-call
  2. Secondary on-call
  3. Tech lead
  4. Platform/SRE

## Incident Severity
- Sev1: total outage or data corruption risk.
- Sev2: partial degradation or persistent high error rate.
- Sev3: minor degradation without business-critical impact.

## Runbook Minimum
- Check `/healthz`, `/readyz`, `/metrics`.
- Check structured logs by `request_id`.
- Check outbox pending/dead volume.
- Validate Kafka broker reachability and producer errors.

## Handover Rules
- Every significant architecture change must include:
  - ADR update
  - risk/trade-off update
  - runbook delta (if operational behavior changed)
