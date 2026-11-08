# Architecture (overview)

Decisions are recorded as ADRs in the docs repository: the docs repository (opening soon). This page is the map.

```
            people                      AI agents (with a person's scoped token)
              │                                      │
        web/ (Next.js)                         mcp/ (MCP server)
              │                                      │
              └──────────────► api/ (Go) ◄───────────┘
                                 │
         ┌───────────────────────┼─────────────────────────┐
         │                       │                         │
   PostgreSQL              Redis (rate limits,       public ledger
   (state, ledger)         real-time, jobs)          (append-only, hash-chained,
                                                     checkpoints published outside)
```

Rules that every component obeys:

1. **One API.** Every use case is an endpoint in `api/openapi/openapi.yaml` before it is anything else. Clients have no privileged path. (ADR-0005)
2. **Secret ballots, public everything else.** Ballots are encrypted on the voter's device and tallied without ever being decrypted individually. Proposals, arguments, delegations, counts and rules are public. (ADR-0003)
3. **The ledger is append-only.** Public events are hash-chained; removals required by law are executed and recorded; nothing is erased in silence. (ADR-0004)
4. **Each person exists once.** Verification levels gate what a person can do; counts are shown per level; verification never multiplies a vote. (ADR-0007)
5. **Agents act in the open.** Actions taken through an agent token are labelled; casting a vote needs the person's own confirmation; limits are per person. (ADR-0005)
6. **Cheap by design.** One VPS plus free tiers must carry the platform until it has real adoption; the budget ceiling is US$100/month. (ADR-0001, ADR-0006)

Domains in `api/internal/` (Phase 2 onwards): `identity`, `scopes` (geography), `proposals`, `debate`, `voting`, `delegation`, `ledger`, `agents`, `notifications`.
