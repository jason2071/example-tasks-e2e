---
description: Scaffold a complete new feature module (entity, repository, service, handler) and wire dependency injection
---

# Workflow: Create Clean Architecture Module

## Objective
Scaffold a complete new feature module strictly following our Go + Fiber v2 + PostgreSQL Clean Architecture pattern.
**Mandatory Data Flow:** `Handler -> Service -> Repository -> Database`

## Steps
1. **Gather Context:** Ask me for the "Module Name" and the "Data Fields (Entity properties)". Wait for my response before proceeding.
2. **Create Domain Entity (`entity.go`):** Create the struct representing the PostgreSQL table. Use appropriate Go types (e.g., `time.Time` for timestamps, `uuid.UUID` for IDs).
3. **Create Repository Layer (`repository.go`):**
   - Define an interface with standard CRUD operations.
   - Create a struct implementation that accepts a PostgreSQL database connection (e.g., `*sql.DB` or your preferred driver/pool) via a constructor function (`NewRepository`).
4. **Create Service Layer (`service.go`):**
   - Define an interface for the business logic.
   - Create a struct implementation that receives the Repository interface via a constructor (`NewService`).
   - Implement the business logic methods, calling the repository.
5. **Create Handler Layer (`handler.go`):**
   - Create a struct that receives the Service interface via a constructor (`NewHandler`).
   - Implement methods that take `*fiber.Ctx` as a parameter.
   - Handle JSON parsing, call the Service, and return appropriate `c.Status().JSON()` responses.

## Strict Constraints
- **Layer Isolation:** Handlers ONLY talk to Services. Services ONLY talk to Repositories. Never skip layers.
- **Framework Agnostic:** Do NOT `import "github.com/gofiber/fiber/v2"` in `entity.go`, `repository.go`, or `service.go`. Fiber belongs exclusively in the Handler layer.
- **Dependency Injection:** Always use constructor functions (`New...`) to inject interfaces.
- **Error Handling:** Bubble up errors from the Repo to the Service, and let the Handler decide the final HTTP status code.

## Expected Output
Generate the 4 files (`entity.go`, `repository.go`, `service.go`, `handler.go`) inside a new folder `internal/<module_name>/`. Provide a brief instruction on how to wire the new handler into the main router.