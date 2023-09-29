# Project-specific variables
include .env

# Directories
MIGRATION_DIR = db/migrations
BINARY = coffeeshop-api


# Docker container configuration
DOCKER_IMAGE = postgres
DOCKER_CONTAINER = coffeeshop-db-container
DB_PORT = 5434


# Go-related variables
GO_CMD = go
GO_BUILD_CMD = $(GO_CMD) build
GO_BUILD_FLAGS = -o $(BINARY)
GO_SRC = cmd/server/main.go


# Makefile targets
.PHONY: help stop_containers create_container create_db start_container create_migration migrate_up migrate_down build run stop


# Display a help message with available targets
help:
	@echo "Coffeeshop API Makefile"
	@echo "-----------------------"
	@echo "Available targets:"
	@echo "  stop_containers     Stop all running Docker containers"
	@echo "  create_container    Create a PostgreSQL Docker container"
	@echo "  create_db           Create a database in the PostgreSQL container"
	@echo "  start_container     Start the PostgreSQL Docker container"
	@echo "  create_migration    Create a new database migration using Goose"
	@echo "  migrate_up          Apply pending database migrations using Goose"
	@echo "  migrate_down        Rollback the last applied database migration using Goose"
	@echo "  build               Build the Go binary for your server"
	@echo "  run                 Build and run the Go server"
	@echo "  stop				 Stop the Go server"


# Stop all running Docker containers
stop_containers:
	@echo "Stopping all running Docker containers..."
	@if [ $$(docker ps -q) ]; then \
		docker stop $$(docker ps -q); \
		echo "Stopped running containers."; \
	else \
		echo "No containers currently running."; \
	fi


# Start the PostgreSQL Docker container
create_container:
	@echo "Creating a PostgreSQL Docker container..."
	docker run --name $(DOCKER_CONTAINER) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d $(DOCKER_IMAGE)
	make create_db


# Start the PostgreSQL Docker container
start_container:
	@echo "Starting the PostgreSQL Docker container..."
	docker start $(DOCKER_CONTAINER)


# Create a new database in the PostgreSQL container
create_db:
	@echo "Creating a new database in the PostgreSQL container..."
	@while ! docker exec $(DB_DOCKER_CONTAINER) pg_isready -h localhost -U $(DB_USER) -q -d $(DB_NAME); do \
		sleep 1; \
		echo "Waiting for PostgreSQL to start..."; \
	done
	sleep 5
	docker exec -it $(DB_DOCKER_CONTAINER) dropdb --username=$(DB_USER) --if-exists $(DB_NAME)
	docker exec -it $(DB_DOCKER_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)


# Create a new database migration using Goose
create_migration:
	@echo "Creating a new database migration using Goose..."
	${HOME}/go/bin/goose -dir=$(MIGRATION_DIR) postgres "user=$(DB_USER) dbname=$(DB_NAME) sslmode=$(DB_SSL_MODE)" create $(DB_NAME) sql


# Apply pending database migrations using Goose
migrate_up:
	@echo "Applying pending database migrations using Goose..."
	goose -dir db/migrations postgres "postgres://${DB_USER}:${DB_PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up


# Rollback the last applied database migration using Goose
migrate_down:
	@echo "Rolling back the last applied database migration using Goose..."
	goose -dir db/migrations postgres "postgres://${DB_USER}:${DB_PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down


# Build the Go binary for your server
build:
	if [ -f "$(BINARY)" ]; then \
		rm ${BINARY}; \
		echo "Deleted Go binary $(BINARY)"; \
	fi
	@echo "Building the Go binary for your server..."
	$(GO_BUILD_CMD) $(GO_BUILD_FLAGS) $(GO_SRC)


# Build and run the Go server
run: build
	@echo "Running the Go server..."
	chmod +x ./$(BINARY)
	./$(BINARY)


# Stop the server
stop:
	@echo "Stopping server..."
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "Server stopped."