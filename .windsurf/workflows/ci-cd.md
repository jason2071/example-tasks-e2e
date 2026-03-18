---
description: Run CI/CD checks for lint, test, build, and deployment readiness
---

# Workflow: CI/CD Validation

## Objective
Verify code changes are ready for merge or release by running quality gates in a predictable order.

## Steps
1. **Confirm Pipeline Scope:** Ask target branch/environment and whether this is PR validation or release validation.
2. **Run Static Checks:** Execute formatting, linting, and vet checks for changed scope.
3. **Run Unit Tests:** Execute unit tests first for fast feedback.
4. **Run Integration/E2E Tests:** Execute API/E2E tests for behavior and contract validation.
5. **Build Artifact:** Run build command and ensure binaries or images are generated successfully.
6. **Check Migration Safety:** Verify pending migrations (if any) have matching up/down scripts and rollback safety.
7. **Collect Results:** Record failed steps with exact command output and affected files.
8. **Publish Readiness Summary:** Mark pipeline as pass/fail and list actions required before merge/deploy.

## Strict Constraints
- Do not skip quality gates unless explicitly requested.
- Keep command sequence deterministic and reproducible.
- Report failures without mutating unrelated project files.

## Expected Output
Provide a CI/CD readiness report with each stage status (lint, unit, e2e, build, migration check) and clear follow-up actions.
