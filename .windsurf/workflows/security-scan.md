---
description: Perform a focused backend security scan and remediation checklist before merge
---

# Workflow: Security Scan

## Objective
Detect common backend security risks early and provide actionable fixes before code is merged or released.

## Steps
1. **Confirm Scan Scope:** Ask which modules/endpoints/PR range should be scanned.
2. **Check Input Validation:** Verify handler input parsing and validation rules for body/query/path values.
3. **Check SQL Safety:** Ensure repository queries use parameterized placeholders and avoid unsafe string interpolation.
4. **Check Error Exposure:** Confirm API responses do not leak DB internals, stack traces, or sensitive details.
5. **Check Auth/Access Boundaries:** Review protected endpoints for missing auth or authorization checks (if applicable).
6. **Check Secret Handling:** Ensure no credentials/tokens are hardcoded or logged.
7. **Run Security Tooling:** Execute available scanners/linters (e.g., `go vet`, `gosec`, dependency checks) where configured.
8. **Summarize Findings:** Report severity-ranked findings with file path, impact, and recommended fix.

## Strict Constraints
- Do not introduce breaking contract changes while remediating security issues unless explicitly requested.
- Prioritize high-severity findings first (injection, auth bypass, sensitive data exposure).
- Keep remediation scoped to security-relevant changes.

## Expected Output
Provide a security scan report with finding severity, evidence, and concrete remediation steps per issue.
