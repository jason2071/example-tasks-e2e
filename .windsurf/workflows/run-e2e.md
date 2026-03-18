---
description: Run end-to-end tests with deterministic setup, execution, and teardown
---

# Workflow: Run E2E Tests

## Objective
Execute end-to-end tests against realistic app wiring and persistence, and report regressions clearly.

## Steps
1. **Confirm Test Scope:** Ask which E2E suite or endpoint group should be executed.
2. **Prepare Environment:** Ensure required env vars, database connection, and migration state are ready.
3. **Reset Test Data:** Clean and seed deterministic test data before running tests.
4. **Run Targeted E2E First:** Execute the smallest relevant suite for quick signal.
5. **Run Broader E2E (If Needed):** Execute full E2E suite after targeted tests pass.
6. **Capture Failures:** For failed tests, collect endpoint, input, expected vs actual, and key logs.
7. **Verify Side Effects:** Confirm expected DB state changes and no unintended mutations.
8. **Summarize Outcome:** Provide pass/fail summary with flaky-risk notes and rerun guidance.

## Strict Constraints
- Keep test runs isolated from production or shared non-test data.
- Do not leave dirty test state after execution.
- Avoid broad reruns before targeted failures are analyzed.

## Expected Output
Provide executed command list, suite-level results, failing case details (if any), and recommended next actions.
