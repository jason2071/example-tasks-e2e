---
description: Add a new API endpoint across handler, service, and repository without breaking existing behavior
---

# Workflow: Add Endpoint

## Objective
Add a new API endpoint end-to-end (handler -> service -> repository) without breaking existing behavior.

## Steps
1. **Confirm Contract:** Define endpoint details first:
   - HTTP method and path.
   - Request payload, path/query params, and response shape.
   - Expected success and error status codes.
2. **Register Route:** Update route registration in bootstrap/router code.
3. **Update Handler:**
   - Add handler method for the new route.
   - Parse and validate request input.
   - Call service only.
4. **Update Service:**
   - Add interface + implementation method.
   - Implement business rules and orchestration.
   - Avoid storage-specific logic.
5. **Update Repository:**
   - Add interface + implementation method for persistence/query.
   - Use parameterized queries (`$1`, `$2`, ...).
   - Select only required fields.
6. **Preserve Compatibility:**
   - Do not rename existing routes, fields, or error envelope unless explicitly requested.
7. **Add Tests:** Add or update tests for success and key failure paths.
8. **Validate and Summarize:** Run focused tests and summarize all changed files.

## Strict Constraints
- Keep layer boundaries strict: Handler -> Service -> Repository only.
- Never introduce direct database calls in handlers or services.
- Preserve existing route contracts unless explicit changes are requested.

## Expected Output
Provide updated handler, service, and repository method signatures/implementations, route registration changes, and a short list of affected tests.
