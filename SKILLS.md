# Go-StudentApp Project Skills

## Purpose
School management backend API built with **Go + Gin + GORM**. Manages schools, students, and courses with nested REST routes.

## Core Architecture
- Entry point: `main.go`
- Modules:
  - `school/`
  - `student/`
  - `course/`
  - `config/`
  - `helper/`
- Layering pattern used throughout:
  1. `controller` (HTTP handlers, DTO mapping)
  2. `service` (business logic, validation)
  3. `repository` (DB CRUD)
  4. `model` (GORM structs)

## Branch
- Current branch in this environment: `fix/student-create-school-validation`
- Upstream: `origin`

## Must-know facts
- Project compiles with Go 1.22+ using local toolchain in this environment.
- DB config is currently in `config/database.go`.
- There are no test files in repository packages.
- Course, student, and school CRUD are implemented without a dedicated migration framework.

## Coding Conventions observed
- Keep existing route compatibility when extending APIs.
- Validate parent-child references in service layer before writes.
- Keep errors mapped through `helper/error.go`.
- Preserve backward-compatible routes/aliases where possible.

## Immediate next improvement ideas
- Externalize DB credentials/env handling.
- Add explicit API docs (README/OpenAPI style endpoint matrix).
- Add schema migrations (course uniqueness/relations).
- Add integration tests for nested route workflows.
