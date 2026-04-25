---
name: owasp-security-review
description: Review changed code for OWASP Top 10 security risks before merge or release.
---

# Skill: OWASP Security Review

## Goal
Run a lightweight but strict security review aligned with OWASP Top 10 before merge/release.
Normative policy source: `.cursor/rules/security-owasp-top10.mdc`.

## Checklist
1. Access control:
- Authorization enforced on protected routes.
- No trust on client-provided role/permission fields.

2. Input and injection:
- All request inputs validated.
- No raw SQL string concatenation.
- Output/error messages do not leak internals.

3. Authentication and secrets:
- API key/token checks applied where required.
- No secrets or credentials in code or logs.

4. Security misconfiguration:
- CORS, rate limit, and body size limits configured centrally.
- Secure defaults in `.env.example`.

5. Logging and monitoring:
- Security-relevant events are logged with `request_id`.
- Logs are structured for centralized aggregation.

6. Data and integrity:
- Migrations are deterministic and reviewed.
- External payload/event data validated before processing.

## Output Format
- Findings first (by severity: high, medium, low).
- Include file references and concrete remediation suggestions.
- If no findings, state explicit pass with residual risk note.
