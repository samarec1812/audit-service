CONFIG_PATH=config/local.yaml
MIGRATIONS_DIR=db/migrations

include .env
POSTGRES_CONNECTION="user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=$(POSTGRES_HOST) port=$(POSTGRES_EXTERNAL_PORT) sslmode=$(POSTGRES_SSLMODE)"
export CONFIG_PATH

.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run -c ./.golangci.yaml ./...

.PHONY: migrate-up
migrate-up:
	goose -dir=$(MIGRATIONS_DIR) postgres $(POSTGRES_CONNECTION) up

.PHONY: migrate-down
migrate-down:
	goose -dir=$(MIGRATIONS_DIR) postgres $(POSTGRES_CONNECTION) down

.PHONY: run
run:
	$(info Run go  project...)
	go run cmd/audit/main.go

