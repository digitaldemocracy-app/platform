# mcp — the official MCP server (planned)

People will give an AI agent a token to operate the platform on their behalf, under their own rules. This directory will hold the official [Model Context Protocol](https://modelcontextprotocol.io/) server that exposes the public API to such agents.

Decisions already taken ([ADR-0005](https://github.com/digitaldemocracy-app/docs/blob/main/adr/0005-api-first-and-ai-agents.md)):

- The MCP server is a **thin client of the public API**: no privileged path, no logic of its own. Anything an agent can do through MCP, a human can do through the same endpoints.
- Written in **Go** with the official MCP Go SDK; transports: Streamable HTTP (hosted at `mcp.digitaldemocracy.app`) and stdio (local).
- Agents authenticate with **agent tokens**: issued by the person, scoped (`read`, `debate`, `propose`, `delegate`, `vote`), expiring, revocable at any time.
- Every action taken through an agent token is **labelled as an agent action** in the ledger and in the UI.
- **Casting a vote always requires the person's own confirmation** (one tap by the owner, not the agent). An agent may prepare a ballot; it cannot cast it. Provisional rule, to be confirmed by product and legal review.
- Rate limits apply **per person**, across all their tokens and sessions, not per token.

Nothing is implemented yet; it starts in Phase 6 of the roadmap, but the API is designed for it from Phase 2.
