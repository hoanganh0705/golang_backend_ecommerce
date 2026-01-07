#name app

APP_NAME = server

dev:
	go run ./cmd/$(APP_NAME)

run:
	docker compose up -d

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

.PHONY: run