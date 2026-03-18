---
description: Execute safe Git workflows for feature work, review readiness, and clean commits
---

# Workflow: Git Workflows

## Objective
Manage day-to-day Git tasks safely, keep history clean, and prepare changes for review without breaking team conventions.

## Steps
1. **Confirm Scope:** Ask which branch/task the user wants to work on and whether there are uncommitted local changes.
2. **Check Workspace State:** Inspect `git status` and identify modified, untracked, or conflicted files.
3. **Review Change Boundaries:** Ensure only task-related files are included; separate unrelated changes.
4. **Stage Intentionally:** Stage files by logical unit (feature, fix, test, docs), not all-at-once by default.
5. **Create Clean Commit Message:** Use clear, scoped commit messages that summarize behavior changes.
6. **Run Pre-Push Validation:** Execute relevant tests/lint/build checks before pushing.
7. **Push and Prepare PR Context:** Push branch and summarize what changed, why, and how it was validated.
8. **Handle Conflicts Safely:** If rebasing/merging causes conflicts, resolve with minimal changes and re-run tests.

## Strict Constraints
- Never use destructive Git commands (`reset --hard`, forced checkout) unless explicitly requested.
- Do not include unrelated file changes in the same commit.
- Preserve existing branch/commit conventions in the repository.

## Expected Output
Provide a concise Git execution summary: branch used, commits created, validation commands run, and PR-ready notes.
