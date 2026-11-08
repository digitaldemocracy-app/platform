# Digital Democracy — platform

[![CI](https://github.com/digitaldemocracy-app/platform/actions/workflows/ci.yml/badge.svg)](https://github.com/digitaldemocracy-app/platform/actions/workflows/ci.yml)
[![License: AGPL-3.0-or-later](https://img.shields.io/badge/license-AGPL--3.0--or--later-blue.svg)](LICENSE)

The source code of the Digital Democracy platform: an open, global platform for constructive debate, secret and verifiable voting, and democratic participation. Public from the first commit.

The [Manifesto](https://github.com/digitaldemocracy-app/manifesto) says *why*. The docs repository says *what* (product, requirements, roadmap) and records *how* (architecture decisions). This repository is the *how*, in code.

_The docs repository (product vision, requirements and architecture decisions) is being finalized and will open soon; until then, references to it below are plain text._

**Status: Phase 1 — Foundations.** There is no product here yet: this is the skeleton (structure, contracts, CI, licence) on which Phase 2 starts. See the roadmap.

## Structure

| Directory | What | Stack |
|---|---|---|
| [`api/`](api/) | Backend: the public API. Every use case exists here first, as a documented endpoint. | Go 1.24+, standard library `net/http`, PostgreSQL, Redis |
| [`web/`](web/) | Frontend: the web app (and, in Phase 1, the landing page). A client of the API like any other. | Next.js 15, React 19, TypeScript, Tailwind CSS 4 |
| [`mcp/`](mcp/) | The official MCP server, so an AI agent can act on a person's behalf under that person's rules. Planned. | Go |
| [`infra/`](infra/) | Local development stack and deployment definitions. | Docker Compose, Caddy, GitHub Actions |
| [`docs/`](docs/) | How to contribute and develop in this repository. Product and architecture docs live in the `docs` repository. | Markdown |

Design rules that shape this layout (each is an ADR in the docs repository):

- **API-first.** The OpenAPI document ([`api/openapi/openapi.yaml`](api/openapi/openapi.yaml)) is the contract; web, mobile and agents are clients. ADR-0005
- **Secret, verifiable, receipt-free voting.** ADR-0003
- **Append-only public ledger; nothing erased in silence.** ADR-0004
- **Identity by verification levels; each person exists once.** ADR-0007
- Stack, monorepo and infrastructure: ADR-0001, ADR-0002, ADR-0006

## Run it

```bash
# backend (Go 1.24+)
cd api && go test ./... && go run ./cmd/api      # http://localhost:8080/healthz

# frontend (Node 22)
cd web && npm ci && npm run dev                  # http://localhost:3000

# databases for local development (optional until Phase 2)
docker compose -f infra/docker-compose.yml up -d
```

More in [`docs/DEVELOPMENT.md`](docs/DEVELOPMENT.md).

## Contributing

Issues and pull requests are welcome, in any language. Read [`docs/CONTRIBUTING.md`](docs/CONTRIBUTING.md) first: Developer Certificate of Origin sign-off, no CLA, no secrets in the repository, and the platform's own principles apply to its code.

Security issues: see [`SECURITY.md`](SECURITY.md). Do not open a public issue for a vulnerability.

## Licence

Code: [GNU AGPL-3.0-or-later](LICENSE). Anyone may run, study, change and redistribute the platform; anyone who runs a modified version as a service must publish their changes. That is the licence that matches *everything is visible*.

Documentation in this repository: [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
