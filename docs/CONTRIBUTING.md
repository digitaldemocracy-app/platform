# Contributing

Thank you. This project is maintained by one volunteer; answers come within seven days. Any language is welcome in issues and pull requests; code, comments and documents are in English.

## Before you start

- **What** the platform should do is decided in the docs repository (product, requirements, ADRs). **How** it is built is decided here. If your change alters a decision, open an issue there first or propose a new ADR.
- Read the [Manifesto](https://github.com/digitaldemocracy-app/manifesto) principles. Code that weakens vote secrecy, verifiability, the ledger, or "each person exists once" will not be merged, however good it is.
- Look for issues labelled `good first issue` and `help wanted`.

## The rules

1. **Licence.** Code is contributed under [AGPL-3.0-or-later](../LICENSE). Documentation under CC BY 4.0. There is no CLA; you keep your copyright.
2. **Developer Certificate of Origin.** Every commit is signed off (`git commit -s`), which means you certify the [DCO](https://developercertificate.org/): you wrote it or have the right to submit it under the licence.
3. **No secrets, ever.** No tokens, keys, passwords, `.env` files or personal data in commits, issues or CI logs. If you commit one by mistake, rotate it immediately and tell us; history is not rewritten to hide it, so rotation is the only fix.
4. **No personal e-mail in commits** if you do not want it public: use GitHub's `noreply` address.
5. **Contract first.** A new or changed endpoint is designed in `api/openapi/openapi.yaml` before it is implemented. The web app and agents are clients; nothing is web-only.
6. **Tests and CI.** Every pull request runs lint, tests and build for `api/`, `web/` and the OpenAPI contract. Keep it green. Bug fixes come with a test that failed before.
7. **Small pull requests**, one topic each, with a description a reviewer can check. Squash merge is used; write a commit title that will make sense in `git log`.
8. **Dependencies** are pinned exactly and updated by Dependabot. Adding one needs a reason in the PR (what it does, why not the standard library, licence compatibility with AGPL).
9. **Neutrality.** No partisan content, imagery or examples in code, fixtures or docs. Test data uses invented proposals.

## Workflow

```bash
git clone https://github.com/digitaldemocracy-app/platform
cd platform
# backend
cd api && go test ./... && cd ..
# frontend
cd web && npm ci && npm run lint && npm test && cd ..
```

Fork, branch from `main`, commit with `-s`, open a pull request. CI runs; a maintainer reviews. Reviews discuss the code, not the person.

## Languages and accessibility

The platform must work for people in ten languages and on cheap phones. Strings in `web/` go through the i18n layer (when it exists; until then keep text in one place), pages work without JavaScript where possible, and every interactive element is reachable by keyboard and screen reader.

## Code style

- Go: `gofmt`, `go vet`, standard library first, `internal/` packages per domain, errors wrapped with context, no global state. Tests with the standard `testing` package.
- TypeScript: strict mode, ESLint config in `web/`, no `any` without a comment, components small and server-first.
- SQL: migrations are append-only files; the ledger tables never get `UPDATE` or `DELETE` grants.
