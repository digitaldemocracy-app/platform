.PHONY: all api web openapi dev-db

all: api web openapi

api:
	cd api && test -z "$$(gofmt -l .)" && go vet ./... && go test -race -count=1 ./...

web:
	cd web && npm ci && npm run lint && npm run typecheck && npm test && npm run build

openapi:
	npx --yes @redocly/cli@1 lint api/openapi/openapi.yaml

dev-db:
	docker compose -f infra/docker-compose.yml up -d
