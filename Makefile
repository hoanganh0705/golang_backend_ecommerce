APP_NAME := server
ENV ?= dev
CONFIG := config/$(ENV).yaml
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:3306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql


.PHONY: dev run up down stop kill upschema downschema resetschema

.PHONY: air

dev:
	go run ./cmd/$(APP_NAME) $(CONFIG)

up:
	docker compose up -d

down:
	docker compose down

stop:
	docker compose stop

kill:
	docker compose kill

upschema:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downschema:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetschema:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset


run: up dev
