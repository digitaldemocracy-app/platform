# api — Digital Democracy backend

Go service exposing the public API. Standard library only for now (`net/http`, Go 1.24+). The OpenAPI document in [`openapi/openapi.yaml`](openapi/openapi.yaml) is the contract: endpoints are designed there first, then implemented.

```bash
cd api
go test ./...
go run ./cmd/api            # listens on :8080 (API_ADDR to change)
curl -s localhost:8080/healthz
curl -s localhost:8080/openapi.yaml | head
```

Layout:

```
cmd/api/          entry point (flags, env, signals)
internal/server/  routing and handlers; no business logic
openapi/          the contract, embedded into the binary and served at /openapi.yaml
```

Planned (see the ADRs in the docs repository): `internal/<domain>/` packages per bounded context (identity, proposals, debate, voting, ledger, delegation, agents), PostgreSQL via `sqlc`, Redis for rate limits and real-time, generated server stubs from the OpenAPI document.
