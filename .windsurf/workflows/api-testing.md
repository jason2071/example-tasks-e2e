---
description: Run API-level validation for Fiber endpoints including success, validation, and error contracts
---

# Workflow: API Testing

## Objective
Validate API behavior at endpoint level for status codes, response schema, validation behavior, and key error paths.

## Steps
1. **Confirm Scope:** Ask which endpoint(s), method(s), and environment should be tested.
2. **Review Contracts:** Read handler DTOs, route registration, and response envelope definitions before testing.
3. **Prepare Data:** Seed required data and clean conflicting state for deterministic results.
4. **Run Success Cases:** Test expected happy paths for each endpoint and verify status/body shape.
5. **Run Validation Cases:** Send invalid payload/query/path values and verify expected error responses.
6. **Run Not-Found/Conflict Cases:** Validate behavior for missing records and duplicate/invalid state transitions.
7. **Check Response Envelope:** Ensure success and error payload formats follow existing project conventions.
8. **Summarize Results:** List passed/failed scenarios with endpoint references and next fixes.

## Strict Constraints
- Do not invent routes, payload fields, or response schemas.
- Keep tests deterministic and isolated from unrelated data.
- Do not expose internal errors in expected API output assertions.

## Expected Output
Provide a concise API test report including tested endpoints, scenario matrix (success/error/validation), and any contract mismatches.
