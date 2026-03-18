---
description: Generate Swagger annotations from handler code for Fiber APIs
---

# Workflow: Generate Swagger

## Objective
Generate or update Swagger annotations from Fiber handler code so API docs match implemented behavior.

## Steps
1. **Review Handlers:** Read target handler methods and existing request/response DTOs.
2. **Write Annotations:** Add or update Swag-style comments above each handler function:
   - Summary, description, tags.
   - Accept/produce content types.
   - Parameters (path/query/body) with correct types and required flags.
   - Success and failure responses with actual schema types.
   - Router path + HTTP method.
3. **Validate Alignment:** Keep comments aligned with real route registration and response envelope.
4. **Check Accuracy:** Do not invent fields or status codes that are not implemented.
5. **Generate Docs:** Run or suggest the project Swagger command (`swag init` or equivalent).
6. **Summarize Coverage:** List documented endpoints and unresolved schema gaps.

## Strict Constraints
- Comments must reflect real handler behavior and real DTOs only.
- Do not invent routes, status codes, or schema fields.
- Keep API response envelope consistent with existing project conventions.

## Expected Output
Provide updated handler comment blocks and the Swagger generation command to run.
