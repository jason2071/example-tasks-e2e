---
description: Generate PostgreSQL up/down migration SQL from the latest entity definition
---

# Workflow: Generate PostgreSQL Migrations

## Objective
Analyze a Go Entity struct and generate standard PostgreSQL up and down migration scripts.

## Steps
1. **Target Entity:** Ask me which Entity/struct I want to generate a migration for. Wait for my response.
2. **Analyze Struct:** Read the specified `entity.go` file. Note all fields, data types, and any specific constraints (e.g., unique, not null).
3. **Map Data Types:** Translate Go types to PostgreSQL types correctly:
   - `string` -> `VARCHAR` or `TEXT`
   - `int` / `int64` -> `INTEGER` or `BIGINT`
   - `time.Time` -> `TIMESTAMP WITH TIME ZONE`
   - `bool` -> `BOOLEAN`
   - `uuid.UUID` -> `UUID`
4. **Generate `Up` Script:** Write the `CREATE TABLE` SQL statement. Include a primary key, standard audit columns (`created_at`, `updated_at`) if applicable, and foreign keys if relations exist.
5. **Generate `Down` Script:** Write the `DROP TABLE` SQL statement to safely rollback the migration.

## Strict Constraints
- **Idempotency:** Use `CREATE TABLE IF NOT EXISTS` and `DROP TABLE IF EXISTS`.
- **Naming Conventions:** Use `snake_case` for table names and column names in PostgreSQL.
- **Syntax:** Strictly adhere to PostgreSQL dialect syntax.

## Expected Output
Output two separate SQL blocks: one for the `.up.sql` file and one for the `.down.sql` file. Suggest a timestamp-based filename format (e.g., `20260318_create_users_table.up.sql`).