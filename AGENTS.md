# AGENTS Instructions for Go Student App

These are hard constraints for any future agent-assisted changes in this repository.

## Non-negotiables
- **No breaking changes by default.**
  - Do not modify API contracts, routes, DTOs, database schema, or validation behavior unless explicitly asked.
- **Preserve architecture.**
  - Keep the existing structure and flow: `router -> controller -> service -> repository -> model`.
- **No broad “auto-refactors.”**
  - Avoid mass edits, automated rename passes, and bulk formatting/rewrite actions unless explicitly requested.
- **No destructive cleanup by default.**
  - Do not delete files or remove behavior without explicit approval.
- **Ask before risky operations.**
  - Before adding/removing dependencies, changing auth, security, database assumptions, or behavior that could break compatibility, confirm first.
- **Backward compatibility first.**
  - If a change can be introduced without breaking compatibility, do that. Keep existing endpoints and behaviors when introducing improvements.

## Scope discipline
- Change only what is required by the user request.
- Keep commit deltas focused and minimal.
- Avoid unrelated feature work and avoid changing file names unless requested.

## Validation expectations
- Run `go test ./...` with the local compatible toolchain before declaring completion.
- If there is no test suite or a command fails, report the exact command and output.

## Security / safety
- Do not introduce hardcoded credentials or secrets.
- Do not expose internal paths/keys in logs or command output.
