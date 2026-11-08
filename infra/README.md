# infra

Local development stack and, later, the deployment definitions. The decision record is [ADR-0006 — Infrastructure and costs](https://github.com/digitaldemocracy-app/docs/blob/main/adr/0006-infrastructure-and-costs.md).

## Local

```bash
docker compose -f infra/docker-compose.yml up -d     # PostgreSQL 17 + Redis 7 on localhost
cp infra/.env.example infra/.env                      # optional overrides; infra/.env is git-ignored
```

## Target topology (Phase 1–3)

```
Internet ──► Cloudflare (DNS, TLS, WAF, CDN, Turnstile)
               ├── digitaldemocracy.app        → Cloudflare Pages (static export of web/)
               └── api.digitaldemocracy.app    → Hetzner VPS (Caddy ─► api binary, PostgreSQL, Redis)
                                                   └── nightly encrypted backups → object storage
GitHub Actions: CI on every PR; deploy = build static binary, ship over SSH, health check, rollback on failure.
```

Everything is in the Founder's or the project's name; nothing inherited from any company. Secrets live in GitHub Environments and a secrets manager, never in this repository. Files here that will appear as the work progresses: `Caddyfile`, `deploy/` (systemd unit, deploy script), `backup/` (pg_dump + age + upload), `cloudflare/` (DNS as code).
