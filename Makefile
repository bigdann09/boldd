include .env

.PHONY: devrun migrate

MIGRATION_PATH := $(shell pwd)/internal/infrastructure/persistence/migrations
DATABASE_URI := postgres://${DATABASE_USER}:${DATABASE_PASS}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=${DATABASE_SSLMODE}

devrun:
	go run cmd/api/main.go

dockerrun:
	docker compose -f docker-compose.dev.yml watch

migrate_create:
	migrate create -ext=sql -dir=${MIGRATION_PATH} -seq ${name}

migrate_up:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URI} up

migrate_down:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URI} down

migrate_fix:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URI} force ${version}

docker-migrate:
	docker run --rm -v ${MIGRATION_PATH}:/migrations migrate/migrate -path=/migrations -database ${DATABASE_URI} up

docker-down:
	docker run --rm -v ${MIGRATION_PATH}:/migrations migrate/migrate -path=/migrations -database ${DATABASE_URI} down

swag-generate:
	swag init -g cmd/api/main.go -o internal/api/docs