---
description: Generate repository and service interface mocks for Go tests
---

# Workflow: Generate Mocks

## Objective
Generate repository and service interface mocks that match existing Go test conventions.

## Steps
1. **Detect Tooling:** Check which mocking tool is already used (`mockery`, `gomock`, or equivalent).
2. **Identify Interfaces:** Select target interfaces (usually repository and service interfaces).
3. **Generate Files:** Generate mocks in the existing mock directory pattern.
4. **Validate Imports:** Ensure package names and import paths match current test conventions.
5. **Preserve Scope:** Do not edit business logic files while generating mocks.
6. **Handle Missing Setup:** If no tool is configured, propose one option with setup steps before generating.
7. **Summarize Results:** List generated files and how to use them in tests.

## Strict Constraints
- Use the existing mocking strategy already present in the codebase.
- Do not change public interfaces only to simplify mock generation.
- Keep generated mocks in test-related directories only.

## Expected Output
Provide the generated mock file paths and the exact command used (or suggested) for regeneration.
