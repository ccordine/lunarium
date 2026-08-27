SYSTEM_DOCKER ?= docker --context default

.PHONY: install build test dev-api dev-web docker-network docker-build docker-up docker-down worknet-publish

install:
	cd web && npm install

build:
	cd web && npm run build
	mkdir -p bin
	GOCACHE=/tmp/lunarium-go-cache go build -o bin/lunarium .

test:
	GOCACHE=/tmp/lunarium-go-cache go test ./...
	cd web && npm run test:navigation
	cd web && npm run test:astronomy
	cd web && npm run build

dev-api:
	GOCACHE=/tmp/lunarium-go-cache go run . -addr 127.0.0.1:8080

dev-web:
	cd web && npm run dev

docker-network:
	@$(SYSTEM_DOCKER) network inspect dev-net >/dev/null 2>&1 || \
		{ echo "System Docker network dev-net is missing; restore the shared development network first." >&2; exit 1; }

docker-build: docker-network
	$(SYSTEM_DOCKER) compose build

docker-up: docker-network
	$(SYSTEM_DOCKER) compose up -d --build

docker-down:
	$(SYSTEM_DOCKER) compose down

worknet-publish:
	curl -fsS -H 'Content-Type: application/json' -X POST \
		"$${WORKNET_API_URL:-http://127.0.0.1:18092}/api/published-apps/from-container" \
		--data @deploy/worknet-published-app.json
