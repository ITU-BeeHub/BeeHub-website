# Build, test and operate the BeeHub website. Run `make help` for a summary.

ENV_FILE   ?= /etc/beehub/env
PREFIX     ?= /opt/beehub
WEB_ROOT   ?= /var/www/beehubapp.com
BACKUP_DIR ?= $(HOME)/backups/beehub
SERVICE    ?= beehub-backend.service
HEALTH_URL ?= http://127.0.0.1:9000/api/health

COMPOSE      := docker compose -f deploy/compose.yaml --env-file $(ENV_FILE)
GO_BUILD     := CGO_ENABLED=0 go build -trimpath -ldflags='-s -w'
TEST_DB_NAME := beehub-test-db
TEST_DB_URL  := postgres://postgres:test@127.0.0.1:55432/postgres?sslmode=disable
# Run psql interactively from a terminal, or read SQL piped to stdin.
DOCKER_EXEC  := docker exec -i$(shell [ -t 0 ] && echo t)

.DEFAULT_GOAL := help
.PHONY: help build backend frontend test test-integration \
        deploy deploy-backend deploy-frontend \
        db-up db-backup db-psql admin-user logs

help: ## Show this help
	@grep -hE '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  %-18s %s\n", $$1, $$2}'

build: backend frontend ## Build the API binaries and the frontend

backend: ## Build the API binaries into build/
	cd backend && $(GO_BUILD) -o ../build/beehub-web ./cmd/website
	cd backend && $(GO_BUILD) -o ../build/beehub-admin-user ./cmd/admin-user

frontend: ## Build the frontend into frontend/dist/
	cd frontend && npm ci --no-audit --no-fund && npm run build

test: ## Run unit tests and linters
	cd backend && go vet ./... && go test -race ./...
	cd frontend && npm run lint

test-integration: ## Run all backend tests against a throwaway PostgreSQL
	docker run -d --rm --name $(TEST_DB_NAME) -e POSTGRES_PASSWORD=test \
		-p 127.0.0.1:55432:5432 postgres:17-bookworm >/dev/null
	until docker exec $(TEST_DB_NAME) pg_isready -h 127.0.0.1 -q; do sleep 1; done
	cd backend && TEST_DATABASE_URL='$(TEST_DB_URL)' go test -race -count=1 ./...; \
		status=$$?; docker rm -f $(TEST_DB_NAME) >/dev/null; exit $$status

deploy: deploy-backend deploy-frontend ## Deploy the API and the frontend

deploy-backend: backend ## Install the API binaries and restart the service
	install -d -m 755 $(PREFIX)/bin
	install -m 755 build/beehub-admin-user $(PREFIX)/bin/beehub-admin-user
	install -m 755 build/beehub-web $(PREFIX)/bin/beehub-web.new
	mv -f $(PREFIX)/bin/beehub-web.new $(PREFIX)/bin/beehub-web
	sudo systemctl restart $(SERVICE)
	@for i in $$(seq 30); do curl -fsS $(HEALTH_URL) >/dev/null 2>&1 && echo "API healthy" && exit 0; sleep 1; done; \
		echo "API did not become healthy, see: journalctl -u $(SERVICE)"; exit 1

deploy-frontend: frontend ## Publish the frontend build to the web root
	rsync -a --delete frontend/dist/ $(WEB_ROOT)/

db-up: ## Start PostgreSQL and wait until it is healthy
	$(COMPOSE) up -d --wait

db-backup: ## Dump the database to BACKUP_DIR (~/backups/beehub)
	mkdir -p $(BACKUP_DIR)
	docker exec beehub-db sh -c 'pg_dump -U "$$POSTGRES_USER" -d "$$POSTGRES_DB" -Fc' \
		> $(BACKUP_DIR)/beehub-$$(date +%Y%m%d-%H%M%S).dump
	@ls -lh $(BACKUP_DIR) | tail -n 3

db-psql: ## Open psql, or run SQL from stdin: make db-psql < file.sql
	$(DOCKER_EXEC) beehub-db sh -c 'psql -v ON_ERROR_STOP=1 -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"'

admin-user: ## Create an admin or reset its password: make admin-user USERNAME=name
	@test -n "$(USERNAME)" || { echo "usage: make admin-user USERNAME=name"; exit 1; }
	set -a && . $(ENV_FILE) && set +a && $(PREFIX)/bin/beehub-admin-user -username '$(USERNAME)'

logs: ## Follow the API logs
	journalctl -u $(SERVICE) -f
