PHONY: dev build tailwind_dev server_dev tailwind_build server_build create_migration migrate_up db_setup seed_db

GOOSE_ENV := GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgresql://postgres:secretpostgres@localhost:5432/splendordb GOOSE_MIGRATION_DIR=migrations
PG_ENV := POSTGRES_URI=postgresql://postgres:secretpostgres@localhost:5432/splendordb
NAME := init_db

dev: tailwind_dev server_dev
	@echo "Starting Dev"
	tailwind_dev
	server_dev

build: tailwind_build server_build
	@echo "Building Prod"
	tailwind_build server_build
	

tailwind_dev:
	@echo "Starting Tailwind watch"
	./tailwindcss -i assets/src/main.css -o assets/dist/main.css --watch

tailwind_build:
	@echo "Building Tailwind"
	./tailwindcss -i assets/src/main.css -o assets/dist/main.css

server_dev:
	@echo "Building Server Dev"
	$(PG_ENV) air

server_build:
	@echo "Building Server Prod"
	go build -o 

create_migration:
	$(GOOSE_ENV) goose create $(NAME) sql

migrate_up:
	$(GOOSE_ENV) goose up

db_setup:
	$(PG_ENV) sqlc generate

seed_db:
	$(PG_ENV) go run cmd/dbseed/main.gochorespot_db
