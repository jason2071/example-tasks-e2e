---
description: Write Go table-driven unit tests for service and handler layers
---

# Workflow: Generate Unit Tests

## Objective
Generate idiomatic Go unit tests for a specific module, focusing on the Service and Handler layers, using table-driven tests and mocking.

## Steps
1. **Target Identification:** Ask me which file or module I want to test. Wait for my response.
2. **Analyze Dependencies:** Identify the interfaces that need to be mocked (e.g., mocking the Repository for the Service test, or mocking the Service for the Handler test).
3. **Generate Mocks:** If mock files don't exist, provide the exact terminal command to generate them (e.g., using `mockery` or `gomock`), or generate manual mock structs if preferred.
4. **Write Service Tests (`service_test.go`):**
   - Use the table-driven test pattern (`[]struct{ name, input, mockSetup, expectedOutput, expectedErr }`).
   - Write test cases for both successful executions and error scenarios.
5. **Write Handler Tests (`handler_test.go`):**
   - Use Fiber's `app.Test(req)` method to test HTTP requests without starting a real server.
   - Assert HTTP status codes and JSON response bodies.

## Strict Constraints
- **No Real Database:** Never connect to a real PostgreSQL database in unit tests. Always use mocked repositories.
- **Coverage:** Ensure edge cases (e.g., not found, validation errors, database connection errors) are included in the test tables.
- **Imports:** Use standard libraries like `testing` and assertion libraries like `github.com/stretchr/testify/assert` if available in the project.

## Expected Output
Provide the full code for the `_test.go` files, formatted cleanly and ready to run via `go test`.