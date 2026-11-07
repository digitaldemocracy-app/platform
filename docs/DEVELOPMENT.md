# Development

## Prerequisites

| Tool | Version | Used by |
|---|---|---|
| Go | 1.24 or newer | `api/`, `mcp/` |
| Node.js | 22 (see `web/.nvmrc`) | `web/` |
| Docker + Compose | any recent | local PostgreSQL and Redis (Phase 2+) |

No other global tools. Linters run through the language toolchains (`gofmt`, `go vet`, `next lint`).

## Everyday commands

```bash
# API
cd api
go run ./cmd/api                 # http://localhost:8080  (API_ADDR to change)
go test -race ./...
gofmt -l . && go vet ./...

# Web
cd web
npm ci
npm run dev                      # http://localhost:3000
npm run lint && npm run typecheck && npm test && npm run build   # same as CI

# Contract
npx --yes @redocly/cli@1 lint api/openapi/openapi.yaml

# Databases
docker compose -f infra/docker-compose.yml up -d
```

`make all` runs the three CI jobs locally.

## Environment variables

Copy `infra/.env.example` to `infra/.env` (git-ignored). The API reads `API_ADDR`, and in Phase 2 `DATABASE_URL`, `REDIS_URL`. Never commit real values.

## Telemetry

Next.js collects anonymous telemetry by default; CI disables it with `NEXT_TELEMETRY_DISABLED=1`. Do the same locally if you prefer: `npx next telemetry disable`.

## Repository conventions

- `main` is protected: pull requests only, CI green, one review, linear history, no force-push.
- Versions follow SemVer; releases are tags `vX.Y.Z` with release notes. There is no release yet.
- Each top-level directory has its own README with its layout.
