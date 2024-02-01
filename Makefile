# To use env variables from local .env file you need to install
# npm install -g dotenv-cli  
include .env
export

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: run
run: # Runs the whole app in docker containers
	docker compose -f ./docker-compose.yaml up -d 

.PHONY: build
build: # Build or rebuild containers
	docker compose -f ./docker-compose.yaml build

.PHONY: down
down: # Stop docker containers
	docker compose -f ./docker-compose.yaml down

.PHONY: migrate postgresql up
migrateup: # migrate postgresql up
	migrate -path ./db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

.PHONY: migrate postgresql down
migratedown: # migrate postgresql down
	migrate -path ./db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down
