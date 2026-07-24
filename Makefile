# Переменные
GO := go
GO_PKG := ./cmd/sharetrip
BINARY_NAME := sharetrip
BUILD_DIR := ./bin
COMPOSE_FILE := ./deploy/docker-compose.yml

# Переменные для миграций
MIGRATE_PATH := ./migration
DB_DSN := postgres://postgres:password@localhost:6543/sharetrip?sslmode=disable

.PHONY: help
help:
	@echo "📦 ShareTrip - Makefile commands"
	@echo ""
	@echo "  make deps          - Install development tools (goimports, golangci-lint, goose)"
	@echo ""
	@echo "  make fmt           - Format Go code"
	@echo "  make lint          - Run golangci-lint"
	@echo "  make test          - Run tests"
	@echo "  make check         - Run all checks (fmt + lint + test)"
	@echo ""
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make e2e           - Run E2E health check (curl)"
	@echo ""
	@echo "  make up            - Start PostgreSQL container"
	@echo "  make down          - Stop PostgreSQL container"
	@echo ""
	@echo "  make migrate-up    - Apply all migrations"
	@echo "  make migrate-down  - Rollback last migration"
	@echo "  make migrate-status- Show migration status"
	@echo ""
	@echo "📖 Examples:"
	@echo "  make deps          # Install tools"
	@echo "  make up            # Start DB"
	@echo "  make migrate-up    # Apply migrations"
	@echo "  make e2e           # Check service health"
	@echo "  make check         # Full CI check"
	@echo "  make build         # Build binary"
	@echo "═══════════════════════════════════════════════════════════"

.PHONY: deps
deps:
	$(GO) install golang.org/x/tools/cmd/goimports@v0.36.0
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.3
	$(GO) install github.com/pressly/goose/v3/cmd/goose@v3.27.0

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	$(GO) test ./...

.PHONY: build
build:
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_PKG)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: run
run:
	$(GO) run $(GO_PKG)

.PHONY: up
up:
	@echo "Starting PostgreSQL container..."
	docker compose -f $(COMPOSE_FILE) up -d
	@echo "✅ PostgreSQL started on port 6543"
	@until docker exec sharetrip-postgres pg_isready -U postgres -h localhost -p 5432; do sleep 1; done
	@echo "✅ Database is ready!"

.PHONY: down
down:
	@echo "Stopping PostgreSQL container..."
	docker compose -f $(COMPOSE_FILE) down
	@echo "✅ Containers stopped"

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations up..."
	goose -dir $(MIGRATE_PATH) postgres "$(DB_DSN)" up
	@echo "✅ Migrations applied"

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back migrations..."
	goose -dir $(MIGRATE_PATH) postgres "$(DB_DSN)" down
	@echo "✅ Migrations rolled back"

.PHONY: migrate-status
migrate-status:
	goose -dir $(MIGRATE_PATH) postgres "$(DB_DSN)" status

.PHONY: check
check: fmt lint test
	@echo "✅ All checks passed!"

.PHONY: e2e
e2e:
	@echo "Running E2E health check..."
	@curl -s http://localhost:8080/api/ready
	@echo ""
	@echo "✅ E2E check passed!"