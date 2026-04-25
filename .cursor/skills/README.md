# Project Skills

These skills are reusable prompt patterns for this repository and follow the Cursor skill directory format.

## Available Skills
- `clean-feature-delivery`
  - Path: `clean-feature-delivery/SKILL.md`
  - Use when implementing a new endpoint/feature or refactoring an endpoint structure while preserving behavior.
- `unit-test-mockery-suite`
  - Path: `unit-test-mockery-suite/SKILL.md`
  - Use when writing or refactoring Go use case unit tests with `mockery` mocks and `testify/suite`.
- `owasp-security-review`
  - Path: `owasp-security-review/SKILL.md`
  - Use when reviewing security risks based on OWASP Top 10.
- `sonarqube-quality-gate`
  - Path: `sonarqube-quality-gate/SKILL.md`
  - Use when checking code quality and maintainability before merge.
- `release-readiness-check`
  - Path: `release-readiness-check/SKILL.md`
  - Use before merge/release/publish to verify architecture, docs, API contract, and operations readiness.

## How To Use
- Mention the skill intent in your prompt (automatic invocation may apply when relevant).
- Use explicit invocation with `/skill-name` when you want a specific skill context.
- Keep output in English for code/comment/docs unless business content requires otherwise.
- Always follow project rules under `.cursor/rules/`.

## Rules vs Skills
- Rules in `.cursor/rules/` define normative project constraints and standards.
- Skills in `.cursor/skills/` define reusable execution workflows and review checklists.
