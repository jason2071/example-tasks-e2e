---
description: Enforce clean architecture boundaries and detect forbidden Fiber imports outside handler layer
---

# Workflow: Enforce Clean Architecture

## Objective
Run a pre-commit architecture review to detect and fix layer-boundary violations.

## Steps
1. **Scan Layers:** Check non-handler layers (`service`, `repository`, `model/entity`, `domain`) for forbidden imports:
   - `github.com/gofiber/fiber/v2`
2. **Report Violations:** If found, list each violation with file path and line reference.
3. **Apply Fixes:** Move transport/web concerns back to handler layer.
4. **Recheck Boundaries:** Confirm layer responsibilities are preserved:
   - Handler: request/response, validation, HTTP mapping.
   - Service: business logic and orchestration.
   - Repository: persistence and SQL.
5. **Check Boundary Leaks:**
   - Service calling HTTP context types.
   - Repository returning HTTP-specific errors.
6. **Run Final Scan:** Ensure zero forbidden imports in non-handler layers.
7. **Summarize Results:** Provide findings and fixes before commit.

## Strict Constraints
- Do not allow Fiber imports outside handler layer.
- Do not move business logic into handler just to satisfy import rules.
- Keep fixes minimal and avoid unrelated refactors.

## Expected Output
Provide a violation report (if any), list corrected files, and confirm final scan result is clean.
