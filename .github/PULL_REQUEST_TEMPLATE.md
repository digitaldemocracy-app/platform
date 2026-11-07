## What

<!-- One paragraph: what changes and why. Link the issue or the ADR. -->

## How to check

<!-- Commands, URLs or screenshots a reviewer can use. -->

## Checklist

- [ ] CI is green (lint, tests, build).
- [ ] No secret, token, key or personal data is in the diff. (`.env*` files are ignored; keep it that way.)
- [ ] New or changed endpoints are documented in `api/openapi/openapi.yaml` first.
- [ ] Anything that touches ballots, identity, tokens or the ledger references the relevant ADR.
- [ ] Commits are signed off (`git commit -s`): I certify the [Developer Certificate of Origin](https://developercertificate.org/).
