APP_NAME := server
ENV ?= dev
CONFIG := config/$(ENV).yaml

.PHONY: dev run up down stop kill

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

run: up dev
