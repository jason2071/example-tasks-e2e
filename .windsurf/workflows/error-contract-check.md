---
description: Validate API error response contracts for consistency across handler, service, and repository flows
---

# Workflow: Error Contract Check

## Objective
Ensure all error responses follow the same API contract and do not leak raw internal details.

## Steps
1. **Identify Contract Source:** Locate canonical error response format in utils/DTO/handler helpers.
2. **Map Error Paths:** List target flows (validation, not found, conflict, internal error) per endpoint.
3. **Review Handler Mapping:** Confirm handlers convert service/repository errors to agreed HTTP status codes.
4. **Review Message Consistency:** Verify error code/message fields remain predictable and stable.
5. **Check Sensitive Data:** Ensure DB/internal stack details are not exposed in API responses.
6. **Run Targeted Tests:** Execute tests for representative error paths and compare actual payloads.
7. **Document Mismatches:** Record file/function and expected vs actual contract mismatches.
8. **Summarize Fix Plan:** Provide prioritized fixes to restore contract consistency.

## Strict Constraints
- Do not invent new error envelope fields without explicit contract updates.
- Preserve backward compatibility for existing consumers whenever possible.
- Keep transport-specific mapping in handler layer, not service/repository.

## Expected Output
Provide an error-contract audit report with compliant paths, mismatches, and concrete remediation actions.
