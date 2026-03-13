---
description: End-to-End (E2E) testing verifies that the entire system works correctly from the user's perspective.
---
# E2E Testing Guide

## Purpose

End-to-End (E2E) testing verifies that the entire system works correctly from the user's perspective.
It tests the full flow:

```
HTTP Request → Handler → Service → Repository → Database → Response
```

E2E tests ensure that integrations between components work as expected.

---

# When to Use E2E Tests

E2E tests should focus on **critical business flows**, not every function.

Good candidates:

* Authentication flow
* Creating important records
* Data pipelines
* External integrations
* Core user journeys

Example:

```
Create Planogram
    → POST /planograms
    → insert into database
    → return created record

Get Planogram
    → GET /planograms/:id
    → return correct data
```

---

# Recommended Test Ratio

```
Unit Tests      ~70%
Integration     ~20%
E2E Tests       ~10%
```

E2E tests are slower and more expensive to maintain.

---

# Stack Assumptions

Backend stack:

* Go
* Fiber v2
* PostgreSQL
* Clean Architecture

Architecture layers:

```
Handler → Service → Repository → Database
```

E2E tests must call the **HTTP API**, not internal functions.

---

# Project Structure

Recommended structure:

```
/test
    /e2e
        setup.go
        teardown.go
        planogram_test.go

/testdata
    create_planogram.json
```

Example:

```
project-root
├─ internal
├─ cmd
├─ test
│   └─ e2e
│       ├─ setup.go
│       ├─ teardown.go
│       └─ planogram_test.go
└─ testdata
```

---

# Test Environment

E2E tests must use a **real database**.

Recommended approach:

* Dedicated test database
* Docker container
* Reset database before tests

Example environment variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=app_test
DB_USER=test
DB_PASSWORD=test
```

---

# Test Server Setup

Create a test server that initializes the full application.

Example:

```go
func SetupTestApp() *fiber.App {
    app := fiber.New()

    db := connectTestDB()

    repo := repository.NewPlanogramRepository(db)
    service := service.NewPlanogramService(repo)
    handler := handler.NewPlanogramHandler(service)

    handler.RegisterRoutes(app)

    return app
}
```

This ensures tests run the **real application stack**.

---

# Example E2E Test

Example: Create Planogram

```go
func TestCreatePlanogram(t *testing.T) {
    app := SetupTestApp()

    body := `{
        "name": "test-planogram"
    }`

    req := httptest.NewRequest(
        http.MethodPost,
        "/planograms",
        strings.NewReader(body),
    )

    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)

    require.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
```

---

# Database Verification

E2E tests should verify database state when necessary.

Example:

```go
var count int

err := db.QueryRow(`
    SELECT count(*)
    FROM planograms
    WHERE name = 'test-planogram'
`).Scan(&count)

require.NoError(t, err)
assert.Equal(t, 1, count)
```

---

# Reset Database Between Tests

Tests must run independently.

Common approach:

```
TRUNCATE TABLE
RESTART IDENTITY
CASCADE
```

Example:

```go
func CleanDB(db *sql.DB) {
    db.Exec(`
        TRUNCATE planograms
        RESTART IDENTITY
        CASCADE
    `)
}
```

Call this before each test.

---

# Using Test Data

Store request and expected payloads in files.

```
/testdata
    create_planogram.json
    create_planogram_expected.json
```

Example loading:

```go
data, _ := os.ReadFile("../testdata/create_planogram.json")
```

Benefits:

* reusable
* readable
* easier to maintain

---

# Running Tests

Run E2E tests:

```
go test ./test/e2e -v
```

Example Makefile:

```
make test-e2e
```

---

# Example E2E Flow

Example full workflow:

```
TestPlanogramFlow

1. Create Planogram
2. Get Planogram
3. Update Planogram
4. Delete Planogram
```

This simulates real user behavior.

---

# Recommended Libraries

HTTP testing:

```
net/http/httptest
```

Assertions:

```
github.com/stretchr/testify
```

Optional tools:

```
testcontainers-go
golang-migrate
```

---

# Best Practices

1. Test only critical flows.
2. Keep tests deterministic.
3. Reset database between tests.
4. Avoid dependencies between tests.
5. Keep test data simple.

---

# Summary

E2E testing validates the system from the user's perspective.

Flow tested:

```
Request → Handler → Service → Repository → Database → Response
```

Focus on critical flows and keep the number of E2E tests small.

Start with a minimal workflow:

```
create → get
```

Then expand as the system grows.
