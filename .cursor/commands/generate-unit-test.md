# Generate Unit Test (Go Backend)

Use this command when generating unit tests for Go backend code in this repository.

## Context Strategy

- Keep context minimal and avoid repeating long test instructions in command prompt.
- Delegate detailed workflow to `@.cursor/skills/unit-test-mockery-suite/SKILL.md`.
- Provide only target scope and behavior expectations to start.

## Required Input

- Target scope: package, file, function, method, or use case.
- Expected behavior: happy path, failure path, edge/validation path.

If target scope is missing, ask for it first.

## Execution Workflow

1. Load `@.cursor/skills/unit-test-mockery-suite/SKILL.md`.
2. Follow the skill checklist for `mockery` + `testify/suite`.
3. Generate only required mocks/tests for the target scope.
4. Keep changes minimal and deterministic.
5. Provide focused `go test` commands for validation (run only if user requests).

## Output

Return results in this structure:

1. Test Scope Understanding
2. Cases Added (happy/failure/edge)
3. Files Changed
4. Validation Commands
5. Risks / Follow-ups
