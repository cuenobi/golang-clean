# Implement Ticket (Jira + FE/Figma Alignment)

Use this command when implementing a Jira ticket end-to-end with backend correctness and frontend alignment.

## Required Input

- Jira ticket key or URL.
- Optional constraints (deadline, scope limits, non-goals).

If ticket key/URL is missing, ask for it first.

## Execution Workflow

1. Read ticket context from Jira via MCP.
   - Fetch ticket summary, description, acceptance criteria, labels, assignee, priority, and status.
   - Read all ticket comments to capture clarifications, edge cases, and late updates.
   - If ticket has attachments/links, inspect relevant ones.

2. Build implementation scope from ticket truth.
   - Extract explicit requirements and non-requirements.
   - List assumptions only when not stated in ticket/comments.
   - Do not implement beyond ticket scope unless required for correctness/safety.

3. Check related cards and FE dependencies.
   - Identify linked/related Jira issues (especially FE tasks).
   - If FE card/design exists, verify expected behavior and UX states against FE source of truth.
   - Prefer checking UI/Figma references when available to prevent backend/FE mismatch.

4. Validate FE/UI expectation before coding.
   - Confirm request/response shape, validation errors, field names, enums/status values, and empty/loading/error states needed by FE.
   - Flag mismatches early and propose minimal compatible adjustments.

5. Implement with project standards.
   - Follow repository rules in `@.cursor/rules/`.
   - Keep architecture boundaries intact.
   - Preserve existing behavior not covered by ticket.
   - Add/update tests for new behavior and critical failure paths.
   - If new endpoint is added, create/update Bruno request in `bruno/golang-clean/...`.
   - If new endpoint is added, add Bruno integration scenario (happy path + failure/validation path).

6. Verify and summarize.
   - Run relevant checks/tests.
   - Verify Bruno collection changes are included for new endpoint work.
   - Verify Bruno integration scenario coverage exists for new endpoint work.
   - Confirm implementation satisfies acceptance criteria.
   - Report any open risks, blockers, or follow-up items.

## Output Format

Return results in this structure:

1. Ticket Understanding
   - Key requirements
   - Important comment clarifications
   - Related FE/design dependencies

2. Implementation Plan
   - Files/modules to change
   - Compatibility considerations with FE/UI

3. Changes Made
   - What was implemented
   - Why this matches ticket + comments + FE/design references
   - Bruno files added/updated (if endpoint added/changed)

4. Validation
   - Tests/checks run
   - Bruno integration scenario coverage for endpoint changes
   - Acceptance criteria status

5. Risks / Follow-ups
   - Remaining uncertainty
   - Suggested next actions

## Guardrails

- Treat Jira ticket + comments as primary requirement source.
- Do not ignore ticket comments; they can override initial description.
- For new endpoint implementation, Bruno request + Bruno integration scenario are mandatory unless user explicitly accepts a documented blocker.
- If FE/Figma references conflict with ticket text, call out conflict explicitly and suggest resolution.
- If required Jira/Figma access is unavailable, stop and ask for access or source links before implementation.
