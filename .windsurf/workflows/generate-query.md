---
description: Generate repository SQL queries with safe parameterized placeholders for PostgreSQL
---

# Workflow: Generate Query

## Objective
Generate readable and safe PostgreSQL queries inside repository code using parameterized placeholders.

## Steps
1. **Confirm Target:** Read the target repository function and confirm input/output contract.
2. **Draft Query:** Write SQL with explicit columns only (avoid `SELECT *`).
3. **Apply Parameters:** Use placeholders (`$1`, `$2`, ...) for all dynamic values.
4. **Review Readability:**
   - Clear `WHERE`, `ORDER BY`, `LIMIT`, and `OFFSET` clauses.
   - Avoid hidden side effects in a single oversized query.
5. **Handle Writes (`INSERT`, `UPDATE`, `DELETE`):**
   - Use `RETURNING` only when needed.
   - Keep transaction boundaries in service layer unless existing repository patterns require otherwise.
6. **Verify Scanning:** Validate scan order and destination fields in Go code.
7. **Check Errors:** Ensure raw internal DB errors are not exposed to API consumers.
8. **Summarize Changes:** Provide final query blocks and mention updated function/file.

## Strict Constraints
- Always use parameterized queries for dynamic values (`$1`, `$2`, ...).
- Keep SQL logic in repository layer only.
- Avoid `SELECT *`; list only required columns.

## Expected Output
Provide final SQL query blocks and the exact repository function(s) updated.
