---
description: Run tests, triage failures, apply focused fixes, and verify regressions are resolved
---

# Workflow: Run Tests and Fix

## Objective
Run relevant test suites, identify root causes of failures, and apply minimal fixes without introducing regressions.

## Steps
1. **Confirm Target Scope:** Ask which feature/module/regression should be validated first.
2. **Run Focused Tests First:** Execute the smallest relevant test command before broad suites.
3. **Capture Failure Details:** Record failing test name, file, assertion mismatch, and stack/error output.
4. **Identify Root Cause:** Trace failing path in handler/service/repository and isolate the defect source.
5. **Apply Minimal Fix:** Change only necessary code while preserving existing contracts and architecture boundaries.
6. **Re-run Targeted Tests:** Confirm the original failure is fixed.
7. **Run Broader Safety Check:** Execute adjacent or full suite as needed to catch regressions.
8. **Summarize Changes:** Report fixed tests, modified files, and remaining risks (if any).

## Strict Constraints
- Do not refactor unrelated code while fixing test failures.
- Maintain existing API contracts unless explicitly requested to change them.
- Prefer targeted validation before running costly full test suites.

## Expected Output
Provide test commands run, before/after results, root-cause summary, and exact files changed to fix failures.
