## Set ENV by running make ENV=local, develop, staging, production
ENV := local
# DB URL is always local
DB_URL := "postgres://postgres:password@localhost:5432/budget_manager?sslmode=disable"
MIGRATE := migrate -path=postgresql/migrations -database "$(DB_URL)"

.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-server 
run-server: ## run the grpc server locally
	./cmd/scripts/run-server-local.sh

.PHONY: run-client 
run-client: ## connect to the local client and local servers
	./cmd/scripts/run-client-local.sh

.PHONY: run-client-prd
run-client-prd: ## connect to the local client and prd server
	./cmd/scripts/run-client-prd.sh

.PHONY: run-auth-client-prd 
run-auth-client-prd: ## connect to the prd auth client
	go run cmd/client/*.go -serverEnv prd -clientName auth -serverTLSPath ssl/ca.crt -serverAddress neelchoudhary.com:1443

.PHONY: run-app-client-prd 
run-app-client-prd: ## connect to the prd app client
	go run cmd/client/*.go -serverEnv prd -clientName app -serverTLSPath ssl/ca.crt -serverAddress neelchoudhary.com:1443


.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-force
migrate-force: ## force migration
	@read -p "Forcing database migration at: " name; \
	$(MIGRATE) force $${name// /_}

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir db/migrations -seq $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-test
migrate-test: ## run migrations, drop migrations, then run migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up
	@$(MIGRATE) down
	@$(MIGRATE) up
	@echo "Success..."