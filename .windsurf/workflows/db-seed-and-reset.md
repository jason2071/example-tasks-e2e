---
description: Prepare deterministic PostgreSQL seed data and safely reset test database state
---

# Workflow: DB Seed and Reset

## Objective
Provide a repeatable process to reset test database state and seed deterministic data for local/E2E testing.

## Steps
1. **Confirm Environment:** Ask which database environment should be targeted (local test, CI test, etc.).
2. **Validate Safety:** Confirm target is non-production before any reset action.
3. **Apply Schema State:** Run required migrations to ensure expected schema version.
4. **Reset Data:** Truncate or clean target tables in dependency-safe order.
5. **Seed Baseline Data:** Insert deterministic records required by tests.
6. **Verify Seed Integrity:** Check row counts and critical reference keys after seeding.
7. **Expose Re-run Command:** Provide a single repeatable command or script path for reruns.
8. **Summarize State:** Report which tables were reset, what data was seeded, and final DB readiness.

## Strict Constraints
- Never run destructive reset actions against production databases.
- Keep seed data deterministic and minimal for test scenarios.
- Use SQL scripts or controlled code paths; avoid manual ad-hoc mutations.

## Expected Output
Provide reset/seed command steps, affected table list, and a readiness confirmation for subsequent tests.
