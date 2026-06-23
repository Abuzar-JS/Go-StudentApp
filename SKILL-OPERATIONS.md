# Go-StudentApp Ops & Contribution Skills

## Local checks
- Run compile/tests:
  ```bash
  ~/.local/go/go1.26.4/bin/go test ./...
  ```
- Note: system default Go in environment may be older; use repo-validated Go binary when needed.

## Common files to inspect
- `main.go`
- `config/database.go`
- `helper/error.go`
- Module route files under `*/controller/router/router.go`
- Service/repository boundaries in each module

## Deployment/runtime notes
- `main.go` initializes DB + validator + route groups and starts Gin server.
- No migration framework is present; any schema change should include a migration strategy outside current run-time path.

## Typical contribution steps
1. Update model/request/service/response/repository consistently.
2. Keep legacy endpoints when modifying behavior.
3. Run compile/test command.
4. Commit changes.
5. Push branch to `origin`.

## Branch state (for this repo)
- Current active branch in this environment: `fix/student-create-school-validation`
