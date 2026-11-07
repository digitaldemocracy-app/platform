# Security policy

A platform that asks people to trust it with their vote has to take reports seriously and answer them in the open, once fixed.

## Reporting a vulnerability

- **Preferred:** GitHub private vulnerability reporting — [Report a vulnerability](https://github.com/digitaldemocracy-app/platform/security/advisories/new). Only maintainers see it.
- Until the project mailbox exists, there is no e-mail channel. When `security@digitaldemocracy.app` is live it will be listed here, with a PGP key.

Please do not open a public issue for a vulnerability, and do not test against production systems with other people's data.

## What to expect

- Acknowledgement within 7 days (one volunteer maintains this project).
- A fix or a mitigation plan within 90 days for confirmed issues; faster for anything that touches ballots, identity or the ledger.
- Public disclosure after the fix, with credit to the reporter if they want it. The advisory is recorded; nothing is erased in silence.

## Scope

Everything in this repository and the services it deploys (`digitaldemocracy.app` and subdomains). Third-party services we use (GitHub, Cloudflare, Hetzner) have their own programmes.

## Supported versions

Only `main` and the latest release. There is no product release yet.
