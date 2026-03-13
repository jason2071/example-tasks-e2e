# Example Tasks API

Example Tasks is a small Go/Fiber service that exposes CRUD APIs for managing tasks stored in PostgreSQL. It uses a clean layering of handler → service → repository plus reusable validation and error helpers.

## Tech Stack

- **Language:** Go 1.24+
- **Framework:** [Fiber v2](https://github.com/gofiber/fiber)
- **Database:** PostgreSQL (schema `example`, table `tasks`)
- **Validation:** go-playground/validator

## Project Structure

```
example-tasks/
├── handler/       # HTTP handlers (Fiber)
├── service/       # Business logic
├── repository/    # Database access layer (sql + pg driver)
├── model/         # Request/response structs + pagination helpers
├── utils/         # Shared error + validation utilities
├── migrations/    # SQL migrations for example.tasks table
├── main.go        # App entry point + Fiber server setup
├── go.mod / sum   # Go modules
└── README.md
```

## Prerequisites

1. Go 1.24 or newer installed locally.
2. PostgreSQL reachable on `localhost` (or update the config in `main.go`).
3. A database named `plms` that contains schema `example` (create it if it does not exist).

> **Database credentials** are currently hard-coded in `main.go` (host/user/password/db). Adjust them before running in other environments or refactor to read from environment variables.

## Database Setup

Run the migration SQL to create the `example.tasks` table, unique constraints, check constraints, and indexes:

```bash
psql -U phumai.ru -d plms -f migrations/0001_create_table_tasks.up.sql
```

To roll back the table:

```bash
psql -U phumai.ru -d plms -f migrations/0001_create_table_tasks.down.sql
```

## Running the Server

```bash
# Install dependencies (first run only)
go mod tidy

# Start the Fiber server on :3000
go run .
```

If everything is configured correctly you should see `Successfully connected to the database!` followed by Fiber listening on port 3000.

## API Reference

### Health Check
- `GET /`
- Returns `Hello, World!` (quick connectivity check)

### Create Task
- `POST /task`
- Body:
  ```json
  {
    "title": "Write documentation",
    "status": "pending",
    "priority": 3
  }
  ```
- Validation: `title` and `status` are required. `priority` (optional) must be 1–5.
- Responses: `201` on success, `409` if a task with the same title already exists.

### List Tasks
- `GET /tasks`
- Query parameters:
  - `cursor` (int64, optional) – pagination cursor (uses task ID for keyset pagination).
  - `size` (int, optional, default 20) – page size.
  - `priority` (int, optional) – filter by priority.
  - `sort_with` (`id`, `priority`, `title`).
  - `sort_by` (`asc`, `desc`).
- Returns a paged payload with `tasks` and `pagination.next_cursor`.

### Get Task by ID
- `GET /task/:id`
- Returns the task object or `404` (`resource not found`).

### Update Task
- `PATCH /task/:id`
- Accepts the same body fields as `POST /task` but all are optional; only provided fields are updated.
- Returns `200` on success, `400` when validation fails, `404` if the record does not exist, `409` for duplicate titles.

### Delete Task
- `DELETE /task/:id`
- Returns `200` when deleted, `404` if the task was not found.

## Error Handling

`utils/errors_response.go` defines reusable `AppError` types (e.g., `E001` = not found, `E003` = duplicate title). Handlers consistently use `utils.HandleError` to translate these into HTTP responses with a JSON body:

```json
{
  "error": "task with the same title already exists",
  "code": "E003",
  "path": "/task/1"
}
```

## Next Steps / Ideas

- Replace hard-coded DB config with environment variables (e.g., `github.com/joho/godotenv`).
- Add automated migration tooling (golang-migrate, atlas, goose) instead of manual `psql` commands.
- Extend pagination to include total count + `has_more` metadata.
- Add authentication/authorization if exposed beyond local usage.
